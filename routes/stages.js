var express = require('express');
var router = express.Router();
var models = require('../models/index');

/* GET StageData listing. */
router.get('/', function(req, res, next) {
  models.StageData.findAll().then(function(stageData){
    res.json(stageData)
  });
});

/* GET single StageData */
router.get('/:userid', function(req, res, next) {
  models.StageData.findOne({
    where: {
      UserId: req.params.userid
    }
  }).then(function(stageData) {
    res.json(stageData);
  });
});

/* POST to StageData listing. */
router.post('/', function(req, res) {
  models.StageData.create(
      req.body
  ).then(function(stageData) {
    res.json(stageData);
  });
});

module.exports = router;
