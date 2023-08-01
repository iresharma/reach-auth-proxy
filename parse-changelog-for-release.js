var parseChangelog = require('changelog-parser')
const fs = require("fs")
parseChangelog('CHANGELOG.md', function (err, result) {
    if (err) throw err

    // changelog object
    const keys = Object.keys(result.versions[1].parsed)
    const parsedData = result.versions[1].parsed
    let str = ""
    if (keys.includes("Added") || keys.includes("Changes")) {
        str += "## Added\n\n- "
        if (keys.includes("Added")) {
            str += parsedData["Added"].join("\n- ")
        }
        if (keys.includes("Changes")) {
            str += parsedData["Changes"].join("\n- ")
        }
    }
    if (keys.includes("Breaking changes")) {
        str += "\n\n## Breaking Changes\n\n- "
        str += parsedData["Breaking changes"].join("\n- ")
    }
    if (keys.includes("Fixes")) {
        str += "\n\n## Fixes\n\n- "
        str += parsedData["Fixes"].join("\n- ")
    }
    if (keys.includes("Migration")) {
        str += "\n\n## Migration\n\n- "
        str += parsedData["Migration"].join("\n- ")
    }
    fs.writeFileSync("release.md", str)
})
