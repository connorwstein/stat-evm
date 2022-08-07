// SPDX-License-Identifier: MIT
pragma solidity >=0.8.0;

interface Median {
    function getMedian(uint256[] memory v1) external view returns (uint256);
}
interface Sampler {
    function sampleRandomNumber(
        uint256 distributionType,
        int256 a,
        int256 b,
        uint256 numSamples
    ) external view returns (int256[] calldata); // TODO: Should this be `memory`?
}
interface Moment {
    function getMoment(uint256 v1, uint256[] memory v2, uint256[] memory v3) external view returns (uint256);
}
interface MatrixMult{
    function matrixMultiply(int256[][] memory a, int256[][] memory b) external view returns (int256[][] memory);
}
interface PricePredict{
    function getPredictPrice(uint256 v1, uint256 v2, uint256 v3) external view returns (uint256);
}

interface IPFSPricePredict {
    function getIPFSPredictPrice(string memory ipfsHash, uint256 v2, uint256 v3) external returns (uint256);
interface IPFSMoment{
    function getMoment(string memory ipfsHash, uint256 moment) external view returns (uint256);
}
interface IPFSFit{
    function getFit(string memory ipfsHash, uint256 fitType) external view returns (int256[][] calldata, int256[] calldata);
}

interface Fit {
    function fit(
        uint256 fitType,
        int256[][] memory X,
        int256[][] memory Y
    ) external view returns (int256[][] calldata); // TODO: Should this be `memory`?
}

contract Test {
    uint256 public med;
    uint256 public moment;
    int256[] public sample;
    int256[][] public product;
    int256[][] public fitted;
    uint256 public prediction;
    uint256 public ipfsPrediction;
    int256[][] public ipfsCoeffs;
    int256[] public ipfsIntercepts;
    uint256 public ipfsMomentRes;
    int256 public predicted;

    event Debug(string message, int256 res);

    Median prec = Median(0x0300000000000000000000000000000000000001);
    Sampler sampler = Sampler(0x0300000000000000000000000000000000000004);
    Moment moment_prec = Moment(0x0300000000000000000000000000000000000006);
    MatrixMult matrixMult = MatrixMult(0x0300000000000000000000000000000000000005);
    IPFSMoment ipfsMoment = IPFSMoment(0x0300000000000000000000000000000000000008);
    IPFSFit ipfsFit =          IPFSFit(0x0300000000000000000000000000000000000009);
    Fit fit = Fit(0x0300000000000000000000000000000000000007);
    PricePredict predict_prec = PricePredict(0x0300000000000000000000000000000000000008);
    IPFSPricePredict ipfs_predict_prec = IPFSPricePredict(0x0300000000000000000000000000000000000010);

    function testMedian(uint256[] memory vals) public {
        med = prec.getMedian(vals);
    }
    
    function testSampler(uint256 distributionType, int256 a, int256 b, uint256 numSamples) public {
        sample = sampler.sampleRandomNumber(distributionType, a, b, numSamples);
    }

    function getLastSample() public view returns (int256[] memory) {
        return sample;
    }

    function testMoment(uint256 v1,uint256[] memory v2, uint256[] memory v3) public {
        moment = moment_prec.getMoment(v1, v2, v3);
    }

    function testMatrixMult(int256[][] memory a, int256[][] memory b) public {
        product = matrixMult.matrixMultiply(a, b);
    }

    function testIPFSMoment(string memory ipfsHash, uint256 moment) public {
        ipfsMomentRes = ipfsMoment.getMoment(ipfsHash, moment);
    }

    function testIPFSFit(string memory ipfsHash, uint256 fitType) public {
        (ipfsCoeffs, ipfsIntercepts) = ipfsFit.getFit(ipfsHash, fitType);
    }

    function getIPFSCoeffs() public view returns (int256[][] memory){
        return ipfsCoeffs;
    }

    function getIPFSIntercepts() public view returns (int256[] memory){
        return ipfsIntercepts;
    }

    function getMatrixMulti() public view returns (int256[][] memory) {
        return product;
    }

    function testFit(uint256 fitType, int256[][] memory X, int256[][] memory Y) public {
        fitted = fit.fit(fitType, X, Y);
    }

    function testPrediction(uint256 steps, uint256 samples, uint256 targetTime) public {
        prediction = predict_prec.getPredictPrice(steps, samples, targetTime);
    }

    function testIPFSPrediction(string memory ipfsHash, uint256 samples, uint256 targetTime) public {
        ipfsPrediction = ipfs_predict_prec.getIPFSPredictPrice(ipfsHash, targetTime, samples);

}
    // Predict a single Y
    function testPrediction(int256[][] memory dependentVars) public  {
       int256[][] memory predictedNoIntercept = matrixMult.matrixMultiply(dependentVars, getIPFSCoeffs());
       // Hack we expect only one val
       int256[] memory intercepts =  getIPFSIntercepts();
       predicted = predictedNoIntercept[0][0]/1e6+intercepts[0]/1e6;
    }

    function getFitted() public view returns (int256[][] memory) {
        return fitted;
    }

    function getPrediction() public view returns (uint256) {
        return ipfsPrediction;
    }
}