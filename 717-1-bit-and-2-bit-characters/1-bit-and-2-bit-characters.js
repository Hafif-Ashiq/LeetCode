/**
 * @param {number[]} bits
 * @return {boolean}
 */
var isOneBitCharacter = function (bits) {
    if (bits.length == 1) {
        return true
    }
    const length = bits.length

    if (bits[length - 1] == 0 && bits[length - 2] == 0) {
        return true
    }

    if (length >= 4) {
        console.log("in-four");

        let forth_bit_zero = false
        if (bits[length - 4] == 0) {
            forth_bit_zero = true
        }
        if (forth_bit_zero && bits[length - 3] == 1) {
            return true
        }
        else if (forth_bit_zero && bits[length - 3] == 0 && bits[length - 2] !== 1) {
            return true
        }

        if (!forth_bit_zero) {
            if (bits[length - 5] == 1 && bits[length - 6] != 1) {
                if (bits[length - 3] == 1)
                    return true
                else if (bits[length - 2] == 1)
                    return false
            }
            else {
                if (bits[length - 2] == 0) {
                    return true
                }
            }

        }
    }
    else {
        if (length % 2 !== 0 && bits[length - 1] == 0 && bits[length - 3] == 1) {
            return true
        }
    }
    return false
};



