const express = require('express')
const router = express.Router()

router.get('/', function (_, res, next) {
  return res.status(200).json({ message: 'UP' })
})

module.exports = router
