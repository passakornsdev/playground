const express = require('express')
const app = express()

const PORT = 8080;

app.get('/v1/api/weather', function (_, res) {
    setTimeout(() => {
        res.status(200).json({temp: Math.floor(Math.random() * 40)})
    }, 500);
})

app.listen(PORT, () => {
    console.log(`Listen on port ${PORT}`)
})