var express = require('express');
var router = express.Router();
var models = require('../models/index');

/* GET TutorialData listing. */
router.get('/', function(req, res, next) {
  models.TutorialData.findAll().then(function(tutorialData){
    res.json(tutorialData)
  });
});

/* GET single TutorialData */
router.get('/:userid', function(req, res, next) {
  models.TutorialData.findOne({
    where: {
      UserID: req.params.userid
    }
  }).then(function(tutorialData) {
    res.json(tutorialData);
  });
});

/* POST to UserProperties listing. */
router.post('/', function(req, res) {
  models.TutorialData.create({
    TutorialDataID: parseInt(req.body.UserPropertyID.toString() + req.body.TutorialID.toString()),
    UserPropertyID: req.body.UserPropertyID,
    TutorialID: req.body.TutorialID,
    Date: req.body.Date
  }).then(function(tutorialData) {
    res.json(tutorialData);
  });
});

module.exports = router;
