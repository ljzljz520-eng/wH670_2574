const fs=require('fs');
const html=fs.readFileSync('index.html','utf8');
fs.writeFileSync('dist.html',html);
console.log('built exam archive frontend');
