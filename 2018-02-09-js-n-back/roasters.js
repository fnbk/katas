


// const reizFenster = 2000;
const reizFenster = 500;
// const reizFolge = "T L H C H S C C Q L C K L H C Q T R R K C H R".split(' ');
const reizFolge = "T L H C H".split(' ');
const eingabe = Array.from(Array(reizFolge.length)).fill('N');

const print = (reiz) => {
    process.stdout.clearLine();
    process.stdout.cursorTo(0);
    process.stdout.write(reiz);
};

// reizFolge.map( reiz => () => print(reiz))
//     .forEach( (printReiz, i) => setTimeout(printReiz, i * reizFenster) );

var i = 0;
reizFolge.map( reiz => () => print(reiz))
    .forEach( (printReiz, j) => {
        setTimeout(() => {
                printReiz();
                i = j;
            }, j * reizFenster)
    });

process.stdin.setRawMode(true);
process.stdin.on('data', () => {
    eingabe[i] = 'J';
});

setTimeout(() => {
    process.stdout.clearLine();
    process.stdout.cursorTo(0);
    console.log(`reizFolge: ${reizFolge.join(' ')}`);
    console.log(`eingabe:   ${eingabe.join(' ')}`);

    process.exit()
    }, reizFolge.length * reizFenster);


// process.stdout.write("hello:\n");
//
//  setTimeout(function() {
// //         process.stdout.clearLine();
// //         process.stdout.cursorTo(0);
//          process.stdout.write("hello");
//      }, 500);
//
// process.stdin.setRawMode(true);
//
// process.stdin.on('data', (d) => {
//     process.stdout.write(`data:${d}`);
//     process.exit();
//     // const chunk = process.stdin.read();
//     // if (chunk !== null) {
//     //     process.stdout.write(`data: ${chunk}`);
//     // }
// });
//
//
// // process.stdin.on( 'data', function( key ){
// //     process.exit();
// // });
