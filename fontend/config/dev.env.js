'use strict'
const merge = require('webpack-merge')
const prodEnv = require('./prod.env')
//本地环境地址
module.exports = merge(prodEnv, {
  NODE_ENV: '"development"',
  ServiceIP:'"http://117.48.209.47:8081/api/obs"'
})


