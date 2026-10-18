idf_component_register(
    SRCS "main.c"
    INCLUDE_DIRS "." "emi"
    REQUIRES esp_wifi esp_event esp_netif nvs_flash driver esp_timer
)

# No libcurl on the ESP32: the generated action only needs the url helpers and
# the websocket client (emi/emiws.h), both plain C.
target_compile_definitions(${COMPONENT_LIB} PRIVATE FETCHX_NO_CURL)
