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
        return file.name.split('.').pop() in validTypes
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
}
