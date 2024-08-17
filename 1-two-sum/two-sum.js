/**
 * @param {number[]} nums
 * @param {number} target
 * @return {number[]}
 */
var twoSum = function (nums, target) {
    let arrayJson = {}
    for (let i = 0; i < nums.length; i++) {
        let remaining = target - nums[i]
        if (arrayJson[remaining] !== undefined) {
            return [arrayJson[remaining], i]
        }
        else {
            arrayJson[nums[i]] = i
        }
    }
};