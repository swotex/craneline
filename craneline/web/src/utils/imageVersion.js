// src/utils/imageVersion.js

export function getDefaultVersion(versions) {
  if (!versions || versions.length === 0) return null

  // priorité 1 : is_latest === true
  const latestFlagged = versions.filter((v) => v.is_latest)

  if (latestFlagged.length === 1) {
    return latestFlagged[0]
  }

  // priorité 2 : s'il y en a 0 ou plusieurs avec is_latest,
  // on prend la plus récente par created_at
  const pool = latestFlagged.length > 0 ? latestFlagged : versions

  return [...pool].sort(
    (a, b) => new Date(b.created_at).getTime() - new Date(a.created_at).getTime()
  )[0]
}