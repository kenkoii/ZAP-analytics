var express = require('express');
var router = express.Router();
var models = require('../models/index');

/* GET UserPurchase listing. */
router.get('/', function(req, res, next) {
  models.UserPurchase.findAll().then(function(userPurchase){
    res.json(userPurchase)
  });
});

/* GET single userPurchase */
router.get('/:userid', function(req, res, next) {
  models.UserPurchase.findOne({
    where: {
      UserId: req.params.userid
    }
  }).then(function(userPurchase) {
    // console.log();
    res.json(userPurchase);
  });
});

/* POST to userPurchase listing. */
router.post('/', function(req, res) {
  models.UserPurchase.create(
    req.body
  ).then(function(userPurchase) {
    res.json(userPurchase);
  });
});

module.exports = router;
