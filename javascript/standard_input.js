const readline = require('readline').createInterface({
    input: process.stdin,
    output: process.stdout,
});

let line = [];
readline.on("line", x => {
    line = x.split(' ').map((el) => parseInt(el));
    readline.close();
}).on("close", function() {
    doRun(line[0], line[1]);
    process.exit();
});

const doRun = (a, b) => {
    console.log(a + b);
}