/**
 * Holds common files methods.
 */
export default class FileHelper {
    /**
     * Check if a file is a file is of valid type.
     * @param {File} file - The target file.
     * @param {[String]} validTypes - The valid types for the file.
     * @returns {boolean}
     */
    static isOfValidType (file, validTypes) {
        const fileExtension = `.${file.name.split('.').pop()}`
        return validTypes.includes(fileExtension)
    }

    /**
     * Gets the name of a file, removing its extension.
     * @param {File} file - The target file.
     * @returns {String}
     */
    static getFilenameWithoutExtension (file) {
        let filename = file.name.substring(file.name.lastIndexOf('/')+1)
        const filenameWithoutExtension = file.name.substring(0, file.name.lastIndexOf('.'))
        return filenameWithoutExtension
    }


    /**
     * Checks if a file is in file list.
     * @param {File} file - The target file.
     * @param {[File]} fileList - The file list.
     * @returns {boolean}
     */
    static isFileInFileList (file, fileList) {
        return Boolean(fileList.find((el) => {return el.name == file.name}))
    }

    /**
     * Sums all file sizes.
     * @param {[File]} files - The array of files.
     * @return {Number}
     */
    static sumFileSizes (files) {
        return files.reduce((total, currentFile) => total + currentFile.size, 0)
    }

    /**
     * Converts the size to megabytes.
     * @param {Number} sizeInBytes - The file size in bytes.
     * @return {Number}
     */
    static convertBytesToMegaBytes (sizeInBytes) {
        const sizeMB = sizeInBytes / Math.pow(1024, 2)
        return Number(parseFloat(sizeMB.toString()).toFixed(2))
    }

    /**
     * Reads a json file.
     * @param {File} file
     * @return {Promise<Object>}
     */
    static async readJsonFile(file) {
        const reader = new FileReader();
        reader.readAsText(file);
        const result = await new Promise((resolve, reject) => {
            reader.onload = function(event) {
                resolve(JSON.parse(reader.result))
            }
        })
        return result
    }
}
