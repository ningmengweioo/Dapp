// SPDX-License-Identifier: GPL-3.0
pragma solidity ^0.8.0;

/**
 * @title 计数器合约
 * @dev 提供简单的计数功能，包括增加、减少和重置计数器
 */
contract Counter {
    // 状态变量，存储计数值
    uint256 public count;

    // 事件，在计数变化时触发
    event CountChanged(uint256 newCount);

    /**
     * @dev 构造函数，初始化计数器为0
     */
    constructor() {
        count = 0;
    }

    /**
     * @dev 将计数器加1
     * @return 增加后的计数值
     */
    function increment() public returns (uint256) {
        count += 1;
        emit CountChanged(count);
        return count;
    }

    /**
     * @dev 将计数器减1
     * @return 减少后的计数值
     */
    function decrement() public returns (uint256) {
        // 确保计数器不会小于0
        require(count > 0, "Count cannot be negative");
        count -= 1;
        emit CountChanged(count);
        return count;
    }

    /**
     * @dev 重置计数器为0
     * @return 重置后的计数值（0）
     */
    function reset() public returns (uint256) {
        count = 0;
        emit CountChanged(count);
        return count;
    }

    /**
     * @dev 获取当前计数值
     * @return 当前计数值
     */
    function getCount() public view returns (uint256) {
        return count;
    }
}