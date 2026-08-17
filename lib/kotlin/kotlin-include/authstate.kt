package emikot

import kotlinx.coroutines.flow.MutableStateFlow
import kotlinx.coroutines.flow.StateFlow
import kotlinx.coroutines.flow.asStateFlow
import kotlinx.serialization.SerialName
import kotlinx.serialization.Serializable

/**
 * AuthState - a small, app-wide "who is signed in" package for Kotlin/Android clients
 * of emi-generated code, playing the same role
 * fireback/ui/packages/auth-client's AuthenticationProvider/useAuthentication do for
 * React: a single place that holds the current session, exposes it reactively, and
 * derives the headers (Authorization/workspace-id/role-id) every generated action's
 * request should carry - see ClientContext.Default in common.kt, and the "Client
 * context & authentication" Kotlin doc page for how the two connect.
 *
 * Unlike the React version there's no `authenticate()`/redirect-to-a-login-page
 * concept baked in - a native app performs its own sign-in flow with whatever
 * network/UI code it wants (an emi action, a WebView, etc.) and just calls
 * `AuthState.setSession(...)` with the result; AuthState only owns *holding* and
 * *persisting* that outcome, not producing it.
 */

/** A workspace the signed-in user belongs to, together with their role in it. */
@Serializable
data class AuthenticationWorkspace(
    @SerialName("workspaceId") val workspaceId: String,
    @SerialName("roleId") val roleId: String? = null,
    @SerialName("name") val name: String? = null,
)

/** Minimal identity info about the signed-in user, as needed for display purposes. */
@Serializable
data class AuthenticationUser(
    @SerialName("uniqueId") val uniqueId: String,
    @SerialName("name") val name: String? = null,
    @SerialName("email") val email: String? = null,
)

/**
 * The authenticated session: a bearer token, who the user is, and which workspaces
 * they can act as - mirrors fireback's AuthenticationSession<TExtra> shape, minus the
 * generic `extra` slot (a Kotlin caller can just define its own richer session type
 * and store it via a custom AuthStateStorage instead of forking this one).
 */
@Serializable
data class AuthenticationSession(
    @SerialName("token") val token: String,
    @SerialName("user") val user: AuthenticationUser? = null,
    @SerialName("workspaces") val workspaces: List<AuthenticationWorkspace> = emptyList(),
    @SerialName("capabilities") val capabilities: List<String> = emptyList(),
)

/**
 * Persists the session/selected-workspace across process restarts. AuthState defaults
 * to InMemoryAuthStateStorage (nothing survives a process restart, but AuthState still
 * works out of the box for tests/non-Android JVM code) - a real Android app calls
 * `AuthState.configure(...)` at startup with its own implementation backed by
 * EncryptedSharedPreferences/DataStore (a bearer token is sensitive - prefer an
 * encrypted store over plain SharedPreferences for saveSession).
 */
interface AuthStateStorage {
    fun loadSession(): AuthenticationSession?
    fun saveSession(session: AuthenticationSession?)
    fun loadSelectedWorkspace(): AuthenticationWorkspace?
    fun saveSelectedWorkspace(workspace: AuthenticationWorkspace?)
}

/** The default AuthStateStorage - nothing persists across process restarts. */
class InMemoryAuthStateStorage : AuthStateStorage {
    private var session: AuthenticationSession? = null
    private var workspace: AuthenticationWorkspace? = null

    override fun loadSession(): AuthenticationSession? = session
    override fun saveSession(session: AuthenticationSession?) {
        this.session = session
    }

    override fun loadSelectedWorkspace(): AuthenticationWorkspace? = workspace
    override fun saveSelectedWorkspace(workspace: AuthenticationWorkspace?) {
        this.workspace = workspace
    }
}

/**
 * The app-wide authentication state. A plain `object` (one instance per process),
 * matching "one AuthenticationProvider per app" on the React side - `session` and
 * `selectedWorkspace` are exposed as StateFlow so both Compose (`.collectAsState()`)
 * and plain coroutine code can observe sign-in/sign-out.
 */
object AuthState {
    private var storage: AuthStateStorage = InMemoryAuthStateStorage()

    private val _session = MutableStateFlow<AuthenticationSession?>(null)
    /** The current session, or null when signed out. */
    val session: StateFlow<AuthenticationSession?> = _session.asStateFlow()

    private val _selectedWorkspace = MutableStateFlow<AuthenticationWorkspace?>(null)
    /**
     * Which of the session's workspaces requests are currently scoped to. Falls back
     * to the session's only workspace when it has exactly one and none has been
     * explicitly selected yet - most users only ever belong to one workspace.
     */
    val selectedWorkspace: StateFlow<AuthenticationWorkspace?> = _selectedWorkspace.asStateFlow()

    /** Equivalent to `session.value != null` - "is anyone currently connected". */
    val isAuthenticated: Boolean
        get() = _session.value != null

    /**
     * Swaps the storage backend and (re)hydrates session/selectedWorkspace from it -
     * call once at app startup, before anything reads `session`/`isAuthenticated`,
     * with a storage implementation backed by real persistence (see AuthStateStorage's
     * doc comment). Safe to skip entirely for tests/short-lived processes, where the
     * default InMemoryAuthStateStorage is enough.
     */
    fun configure(storage: AuthStateStorage) {
        this.storage = storage
        val restored = storage.loadSession()
        _session.value = restored
        _selectedWorkspace.value = storage.loadSelectedWorkspace() ?: defaultWorkspace(restored)
    }

    /**
     * Stores a freshly-obtained session (e.g. after a sign-in action's response) and
     * persists it. Pass null to sign out (equivalent to calling signOut()).
     */
    fun setSession(session: AuthenticationSession?) {
        storage.saveSession(session)
        _session.value = session
        if (session == null) {
            selectWorkspace(null)
        } else {
            _selectedWorkspace.value = storage.loadSelectedWorkspace() ?: defaultWorkspace(session)
        }
    }

    /** Selects (and persists) which workspace/role subsequent requests should be scoped to. */
    fun selectWorkspace(workspace: AuthenticationWorkspace?) {
        storage.saveSelectedWorkspace(workspace)
        _selectedWorkspace.value = workspace
    }

    /** Clears the session and selected workspace, both in memory and in storage. Local only - does not call the backend. */
    fun signOut() {
        setSession(null)
    }

    /**
     * `{ "Authorization": ..., "workspace-id": ..., "role-id": ... }` derived from the
     * current session/selected workspace, ready to merge into a ClientContext's
     * defaultHeaders (see the "Client context & authentication" Kotlin doc page) -
     * fields with no value (e.g. no workspace selected yet, or signed out) are simply
     * omitted rather than sent as an empty/"null" string.
     */
    fun headers(): Map<String, String> {
        val result = mutableMapOf<String, String>()
        _session.value?.token?.let { result["Authorization"] = it }
        _selectedWorkspace.value?.let { ws ->
            result["workspace-id"] = ws.workspaceId
            ws.roleId?.let { result["role-id"] = it }
        }
        return result
    }

    /**
     * Re-checks whether the current session is still valid, via a caller-supplied
     * check (e.g. an emi action call to a whoami/verify endpoint). Resolves true if
     * the session is (still) considered valid, false if it was found invalid and
     * cleared as a result. Resolves false immediately, without calling `check`, when
     * there's no session to begin with. A thrown exception from `check` (a network
     * blip, not an explicit "invalid") does not sign the user out, matching
     * fireback's AuthenticationProvider.checkValidity behavior.
     */
    suspend fun checkValidity(check: suspend (AuthenticationSession) -> AuthenticationSession?): Boolean {
        val current = _session.value ?: return false
        return try {
            val refreshed = check(current)
            if (refreshed != null) {
                setSession(refreshed)
                true
            } else {
                setSession(null)
                false
            }
        } catch (e: Exception) {
            true
        }
    }

    private fun defaultWorkspace(session: AuthenticationSession?): AuthenticationWorkspace? {
        return if (session?.workspaces?.size == 1) session.workspaces[0] else null
    }
}
