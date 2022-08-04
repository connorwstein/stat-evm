// SPDX-License-Identifier: MIT
pragma solidity >=0.8.0;

interface Median {
    function getMedian(uint256 v1, uint256 v2, uint256 v3) external view returns (uint256);
}
interface Sampler {
    function getSample(uint256 v1, uint256 v2) external view returns (uint256);
}

contract Test {
    uint256 public last = 0;
    uint256 public sample = 3;

    event Debug(string message, bytes32 res);

    Median prec = Median(0x0300000000000000000000000000000000000001);
    Sampler sampler = Sampler(0x0300000000000000000000000000000000000004);

    function testMedian(uint256 v1, uint256 v2, uint256 v3) public {
        last = prec.getMedian(v1, v2, v3);
    }
    
    function testSampler(uint256 v1, uint256 v2) public {
        sample = sampler.getSample(v1, v2);
    }
}