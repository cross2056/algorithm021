学习笔记

1.two-sum

    key: 数组题, 用哈希表来实现o(n)算法复杂度

21.merge-two-sorted-lists

    key: 链表题，添加虚拟头部指针, 遍历比较两个链表头节点并添加，添加后虚拟指针移动到添加的节点位置

26.remove-duplicates-from-sorted-array

    key: 数组题，单循环双指针法

42.trapping-rain-water
    key: 单调栈题，遍历柱子heights按大小单调递减把index入栈，当遍历到的当前柱子高度大于栈顶index对应柱子高度，^栈顶出栈，然后计算面积sum = h*w，h为min(新栈顶index对应柱子高度,当前待入栈index对应高度)-出栈的对应柱子高度，w为待入栈元素index-新栈顶index-1，计算当前面积sum后再次拿当前柱子高度与新栈顶柱子高度比较，大于新栈顶的话重复上面计算流程，并累计sum^，小于的话即入栈完毕。然后继续遍历下一根柱子，直到柱子遍历完毕，返回sum

66.plus-one

    key: 数组题，最高位计算结果大于10需要扩展数组

88.merge-sorted-array

    key: 数组题，三指针法，指针1指向数组1最后数据位置，指针2指向数组2最后数据位置，指针3指向数组1最后位置，如果数组1数据比数组2少，数组2剩余数据需要结合指针2完成拷贝

189.rotate-array

    key: 数组题，暴力破解：最后k位前移, 或者脑筋急转弯：1.反转整个数组，2.反转前k个数，3.反转后k个数

283.move-zeroes

    key: 数组题，两个循环一个遍历移动数字，另一个剩余补零

641.design-circular-deque

    key: 数组题，front和rear指针，front指针指向下一个frontinsert位置的前一个位置，rear指针指向下一个rearinsert位置，数组capacity申请加k+1，front先-1再赋值，rear先赋值再加1，取rear值得先-1再取值