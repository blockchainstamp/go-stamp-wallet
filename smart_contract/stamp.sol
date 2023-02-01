// SPDX-License-Identifier: GPL-3.0

pragma solidity >=0.8.1;

contract Owner {
    address public owner;
    mapping(address => bool) public admins;

    event OwnerSet(address indexed oldOwner, address indexed newOwner);
    event AdminOP(address indexed newAdmin, bool add);
    modifier isOwner() {
        require(msg.sender == owner, "Caller is not owner");
        _;
    }
    modifier isAdmin() {
        require(admins[msg.sender], "Caller is not admin");
        _;
    }

    constructor() {
        owner = msg.sender;
        admins[msg.sender] = true;
    }

    function changeOwner(address newOwner) public isOwner {
        emit OwnerSet(owner, newOwner);
        owner = newOwner;
    }

    function adminOp(address addr, bool isAdd) public isOwner {
        if (isAdd) {
            admins[addr] = true;
        } else {
            delete admins[addr];
        }
        emit AdminOP(addr, isAdd);
    }
}

contract Stamp is Owner {
    struct Record {
        uint256 value;
        uint256 nonce;
    }
    struct Config {
        string MailBoxName;
        bool IsConsummable;
    }

    mapping(address => Record) private _balances;

    uint256 public totalSupply;
    Config public conf;

    event Transfer(address indexed to, uint256 value);
    event Revoke(address indexed to, uint256 value);

    function transfer(address to, uint256 value) public isAdmin {
        require(to != address(0));
        _balances[to].value = _balances[to].value + value;
        _balances[to].nonce = _balances[to].nonce + 1;
        totalSupply = totalSupply + value;
        emit Transfer(to, value);
    }

    function revoke(address to, uint256 value) public isAdmin {
        require(to != address(0));
        require(_balances[to].value > value);
        _balances[to].value = _balances[to].value - value;
        _balances[to].nonce = _balances[to].nonce + 1;
        totalSupply = totalSupply - value;
        emit Revoke(to, value);
    }

    function BalanceOf(address user) public view returns (Record memory) {
        return _balances[user];
    }

    constructor(string memory mailboxName, bool consummable) {
        conf = Config(mailboxName, consummable);
    }
}
