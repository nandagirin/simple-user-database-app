const express = require('express');
const router = express.Router();
const service = require('../services');

router.post('/', function(req, res, next) {
    return service.authService.generateToken(req, res, next);
})

module.exports = router;
