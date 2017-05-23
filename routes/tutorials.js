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
      UserId: req.params.userid
    }
  }).then(function(tutorialData) {
    res.json(tutorialData);
  });
});

/* POST to UserProperties listing. */
router.post('/', function(req, res) {
  models.TutorialData.create({
    tutorialDataId: parseInt(req.body.userId.toString() + req.body.tutorialId.toString()),
    userId: req.body.userId,
    tutorialId: req.body.tutorialId,
    date: req.body.date
  }).then(function(tutorialData) {
    res.json(tutorialData);
  });
});

module.exports = router;
