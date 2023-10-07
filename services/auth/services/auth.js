const jsonwebtoken = require('jsonwebtoken');

const JWT_SECRET = process.env.JWT_SECRET;
const ADMIN_PASS = process.env.ADMIN_PASS;

const generateToken = (req, res, next) => {
    const { username, password } = req.body;
    if (username === "admin" && password === ADMIN_PASS ) {
        return res.json({
            token: jsonwebtoken.sign({user: "admin"}, JWT_SECRET),
        });
    }

    return res.status(401).json({ message: "Bad credentials" });
}

module.exports = { generateToken };
