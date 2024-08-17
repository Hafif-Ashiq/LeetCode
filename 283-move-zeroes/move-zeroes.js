var moveZeroes = function (nums) {
    var i = 0
    const length = nums.length
    while (i < length) {
        if (nums[i] == 0) {
            console.log(i);
            for (let j = i; j < length; j++) {
                if (nums[j] != 0) {
                    var temp = nums[i]
                    nums[i] = nums[j]
                    nums[j] = temp
                    break
                }
            }

        }

        i++

    }
    return nums
};