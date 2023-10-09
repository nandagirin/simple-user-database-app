const express = require('express')
const logger = require('morgan')
const router = require('./routes')
const app = express()

app.use(logger('dev'))
app.use(express.json())
app.use(express.urlencoded({ extended: false }))

app.use('/auth', router.authRouter)
app.use('/healthz', router.healthCheckRouter)

module.exports = app
