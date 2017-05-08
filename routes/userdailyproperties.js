var express = require('express');
var router = express.Router();
var models = require('../models/index');

/* GET UserDailyProperties listing. */
router.get('/', function(req, res, next) {
  models.UserDailyProperty.findAll().then(function(userDailyProperty){
    res.json(userDailyProperty)
  });
});

/* GET single userproperty */
router.get('/:userid', function(req, res, next) {
  models.UserDailyProperty.findOne({
    where: {
      UserID: req.params.userid
    }
  }).then(function(userDailyProperty) {
    // console.log();
    res.json(userDailyProperty);
  });
});

/* POST to UserProperties listing. */
router.post('/', function(req, res) {
  models.UserDailyProperty.create(
    req.body
  ).then(function(userDailyProperty) {
    res.json(userDailyProperty);
  });
});

module.exports = router;
