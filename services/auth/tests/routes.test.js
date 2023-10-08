const request = require('supertest');
const app = require('../app.js')
const ADMIN_PASS = process.env.ADMIN_PASS;

describe('Auth endpoint', () => {
  it('should generate valid token', async () => {
    const res = await request(app)
      .post('/auth')
      .send({
        username: "admin",
        password: ADMIN_PASS,
      });
    expect(res.statusCode).toEqual(200);
    expect(res.body).toHaveProperty('token');
  });

  it('should not generate valid token and return 401', async () => {
    const res = await request(app)
      .post('/auth')
      .send({
        username: "admin",
        password: "invalid_pass",
      });
    expect(res.statusCode).toEqual(401);
    expect(res.body).not.toHaveProperty('token');
  });
});

describe('Health endpoint', () => {
  it('should return 200 response code', async () => {
    const res = await request(app)
      .get('/healthz');
    expect(res.statusCode).toEqual(200);
  });
});
