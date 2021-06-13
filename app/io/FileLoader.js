const fs = require('fs');

module.exports = class FileLoader {
    
    //Method used to load the file from the selected path
    static loadFile(path) {

        try {
            let buffer = fs.readFileSync(path)
            return {
                success: true,
                name: this.getFileNameFromPath(path),
                data: buffer
            }
        } catch (_) {
            return {
                success: false,
                error: "File not found"
            }
        }
    }

    //method used to get the file name from the path
    static getFileNameFromPath(path) {
        let pathComponents = path.split("/");
        return pathComponents[pathComponents.length - 1];

    }
}