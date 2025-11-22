const express =require('express')
const expressProxy=require('express-http-proxy');

const app=express()

app.get('/',(req,res)=>{
    res.send('Welcome to the Gateway');
});

app.use('/service1',expressProxy('http://localhost:3001'));
app.use('/service2',expressProxy('http://localhost:3002'));

app.listen(3000,()=>{
    console.log('Gateway is running on port 3000');
});

