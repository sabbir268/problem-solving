"use strict";

const fs = require("fs");

process.stdin.resume();
process.stdin.setEncoding("utf-8");

let inputString = "";
let currentLine = 0;

process.stdin.on("data", (inputStdin) => {
    inputString += inputStdin;
});

process.stdin.on("end", (_) => {
    inputString = inputString
        .replace(/\s*$/, "")
        .split("\n")
        .map((str) => str.replace(/\s*$/, ""));

    main();
});

function readLine() {
    return inputString[currentLine++];
}
// https://www.hackerrank.com/challenges/jumping-on-the-clouds-revisited/problem
// Complete the jumpingOnClouds function below.
function jumpingOnClouds(c, k) {
    var E = 100;
    var current_cloud = 0;
    var n = c.length;
    while (true) {
        current_cloud += k;
        if (current_cloud >= n) {
            current_cloud -= n;
        }
        E -= 1;
        if (c[current_cloud] == 1) {
            E -= 2;
        }
        if (current_cloud == 0) {
            console.log(E);
            break;
        }
    }

    return E;
}

function main() {
    const ws = fs.createWriteStream(process.env.OUTPUT_PATH);

    const nk = readLine().split(" ");

    const n = parseInt(nk[0], 10);

    const k = parseInt(nk[1], 10);

    const c = readLine()
        .split(" ")
        .map((cTemp) => parseInt(cTemp, 10));

    let result = jumpingOnClouds(c, k);

    ws.write(result + "\n");

    ws.end();
}