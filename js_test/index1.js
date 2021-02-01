var str = "";
for (var i = 1; i <= 9; i++) {
    for (var j = 1; j <= i; j++) {
        str += i + "x" + j + "=" + i * j + "\t";
    }
    str += "\n";
}
console.log(str);

// for (var i = 0, len = arr.length; i < len; i++) {
//   str += 1;
// }
// var nam = prompt('请输入你的名字');
// var sex = prompt('请输入你的性别');
// var age = prompt('请输入你的年龄');
// alert('你的名子是' + nam);
// alert('你的性别是' + sex);
// alert('你的年龄是' + age);
// var age = 21;
// console.log('wo');
// console.log((age.toString()).length);
// // console.log(typeof age);
// console.log(parseInt('3.14'));
// console.log(parseFloat('3.14'));
// console.log(parseInt(3.14));

// console.log(age++);
// console.log(age);
// var year = prompt('请输入年份');
// if ((year % 4 == 0 && year % 100 != 0) || year % 400 == 0) {
//     alert('闰年');

// } else {
//     alert('平年');
// }
// var num = prompt('hangshu');
// var str = '';
// for (var i = 1; i <= num; i++) {
//     for (var j = 1; j <= num; j++) {
//         str += '*';
//     }
//     str += '\n';
// }
// console.log(str);
// var num = prompt('hangshu');
// var str = '';
// for (var i = 1; i <= num; i++) {
//     for (var j = 1; j <= num - i; j++) {
//         str += '*';
//     }
//     str += '\n';
// }
// console.log(str);
// var arr = [2, 3, 0, 1, 23, 0, 56, 0, 9];
// var newArr = [];
// for (var i = 0, len = arr.length; i < len; i++) {
//     if (arr[i] != 0) {
//         newArr[newArr.length] = arr[i];
//     }
// }
// console.log(newArr);
// var arr = [5, 4, 3, 2, 1];
// var temp;
// for (var j = 0; j < arr.length; j++) {
//     for (var i = 0, len = arr.length - 1; i < len; i++) {
//         if (arr[i] > arr[i + 1]) {
//             temp = arr[i];
//              arr[i] = arr[i + 1];
//             arr[i + 1] = temp;
//         }
//     }
// }
// console.log(arr);
