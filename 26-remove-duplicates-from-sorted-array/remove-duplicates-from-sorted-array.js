/**
 * @param {number[]} nums
 * @return {number}
 */
var removeDuplicates = function (nums) {
    const leng = nums.length
    const n = []
    for (let i = 0; i < leng; i++) {
        if (!n.includes(nums[0])) {
            n.push(nums[0])
        }
        nums.shift()
    }
    n.forEach(num => {
        console.log(num);

        nums.push(num)
    })

    return nums.length

};
