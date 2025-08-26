export function applyFlags(cmd, flags) {
  for (const f of flags) {
    switch (f.Type) {
      case "string":
        cmd.option(`--${f.Name} <value>`, f.Usage, f.Default || undefined);
        break;
      case "bool":
        cmd.option(`--${f.Name}`, f.Usage);
        break;
      case "int":
        cmd.option(`--${f.Name} <number>`, f.Usage, parseInt);
        break;
    }
  }

  // common flags
  cmd.requiredOption("--path <file>", "Path of the file on disk");
  cmd.option("--output <dir>", "Directory to write generated files");
  cmd.option(
    "--tags <tags>",
    "Comma separated tags (e.g. angular,typescript,react,axios,nestjs)"
  );
}
