import { ethers } from "hardhat";

async function main() {
  const Lock = await ethers.getContractFactory("Lock");
  const lock = await Lock.attach("0x1f7262D26e5cC6c92F38684f23A69D7172954f0B");

  await lock.withdraw();
}

// We recommend this pattern to be able to use async/await everywhere
// and properly handle errors.
main().catch((error) => {
  console.error(error);
  process.exitCode = 1;
});
