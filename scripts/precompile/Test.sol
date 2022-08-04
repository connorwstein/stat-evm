// SPDX-License-Identifier: MIT
pragma solidity >=0.8.0;

interface Median {
    function getMedian(uint256 v1, uint256 v2, uint256 v3) external view returns (uint256);
}
interface Sampler {
    function getSample(uint256 v1, uint256 v2) external view returns (uint256);
}
interface Simple {
    function getSimple(uint256 v1) external view returns (uint256);
}

contract Test {
    uint256 public last = 0;
    uint256 public sample = 1;
    uint256 public lastSimple = 2;

    event Debug(string message, uint256 res);

    Median prec = Median(0x0300000000000000000000000000000000000001);
    Sampler sampler = Sampler(0x0300000000000000000000000000000000000004);
    Simple simple = Simple(0x0300000000000000000000000000000000000005);

    function testMedian(uint256 v1, uint256 v2, uint256 v3) public {
        emit Debug("Part1", v1);
        last = prec.getMedian(v1, v2, v3);
        emit Debug("Part2", last);
    }
    
    function testSampler(uint256 v1, uint256 v2) public {
        emit Debug("Part1", v1);
        sample = sampler.getSample(v1, v2);
        emit Debug("Part2", sample);
    }

    function testSimple(uint256 v1) public {
        emit Debug("Part1", v1);
        lastSimple = simple.getSimple(v1);
        emit Debug("Part2", lastSimple);
    }
}