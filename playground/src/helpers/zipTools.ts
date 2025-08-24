import JSZip from "jszip";
import { saveAs } from "file-saver";

export type VirtualFile = {
  Name: string;
  MimeType: string;
  Location: string; // e.g. "src/components/generated"
  ActualScript: string;
  Extension: string; // may be empty
};

/**
 * Build a zip file from VirtualFile[]
 */
export async function buildZip(files: VirtualFile[]): Promise<Blob> {
  const zip = new JSZip();

  for (const f of files) {
    // build filename
    const filename = f.Extension ? `${f.Name}${f.Extension || ""}` : f.Name;
    const path = f.Location ? `${f.Location}/${filename}` : filename;

    zip.file(path, f.ActualScript, { binary: false });
  }

  return await zip.generateAsync({ type: "blob" });
}

/**
 * Download the zip file in browser
 */
export async function downloadZip(
  files: VirtualFile[],
  filename = "sdk-output.zip"
) {
  const blob = await buildZip(files);
  saveAs(blob, filename);
}
