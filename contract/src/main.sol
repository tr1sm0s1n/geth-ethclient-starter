// SPDX-License-Identifier: MIT
pragma solidity 0.8.28;

contract Cert {
    event Issued(string course, uint256 id, string date);

    struct Certificate {
        string name;
        string course;
        string grade;
        string date;
    }

    address admin;
    mapping(uint256 => Certificate) certificates;

    constructor() {
        admin = msg.sender;
    }

    function issue_certificate(
        uint256 _id,
        string memory _name,
        string memory _course,
        string memory _grade,
        string memory _date
    ) public {
        require(admin == msg.sender, "Access Denied");
        certificates[_id] = Certificate(_name, _course, _grade, _date);
        emit Issued(_course, _id, _grade);
    }

    function fetch_certificate(
        uint256 _id
    ) public view returns (Certificate memory) {
        return certificates[_id];
    }
}
