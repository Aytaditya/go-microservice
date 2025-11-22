const express=require('express');

const app=express();

app.get('/',(req,res)=>{
    res.send('Welcome to Service 2');
})

app.listen(3002,()=>{
    console.log('Service 2 is running on port 3002');
})