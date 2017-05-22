var express = require('express');
var router = express.Router();
var models = require('../models/index');

/* GET UserProperties listing. */
router.get('/', function(req, res, next) {
  models.UserProperty.findAll().then(function(userProperty){
    res.json(userProperty)
  });
});

/* GET single userproperty */
router.get('/:userid', function(req, res, next) {
  models.UserProperty.findOne({
    where: {
      UserID: req.params.userid
    }
  }).then(function(userProperty) {
    res.json(userProperty);
  });
});

/* POST to UserProperties listing. */
router.post('/', function(req, res) {
  models.UserProperty.upsert(
    req.body,
      {
      where: {
        UserPropertyID: req.body.UserPropertyID
      }
  }).then(function(userProperty) {
    res.json(userProperty);
  });
  
});

module.exports = router;


