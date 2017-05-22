var express = require('express');
var router = express.Router();
var models = require('../models/index');
var sequelize = models.sequelize;

/* GET home page. */
router.get('/', function(req, res, next) {
  res.render('index', { title: 'FreCre Analytics' });
});

router.post('/', function(req, res, next) {
  const items = req.body;
  let count = 0;
  let promises = [];
  for(var i = 0; i < items.length; i++){
    var item = items[i];
    switch(item.className){
      case 'Stage':
        promises.push(insertStage(item.params)
          .then((res)=>{
          count++;
        }).catch((err)=>{
          console.log(err);
        }));
        break;
      case 'UserProperty':
        promises.push(insertUserProperty(item.params)
          .then((res)=>{
          count++;
        }).catch((err)=>{
          console.log(err);
        }));
        break;
      case 'UserPurchase':
        promises.push(insertUserPurchase(item.params)
          .then((res)=>{
          count++;
        }).catch((err)=>{
          console.log(err);
        }));
        break;
      case 'UserDailyProperty':
        promises.push(insertUserDailyProperty(item.params)
          .then((res)=>{
          count++;
          console.log(count);
        }).catch((err)=>{
          console.log(err);
        }));
        break;
      case 'Tutorial':
        promises.push(insertTutorial(item.params)
          .then((res)=>{
          count++;
        }).catch((err)=>{
          console.log(err);
        }));
        break;
    }
  }
  Promise.all(promises)
    .then(() => {
      res.json(count);
    });
});

function insertStage(item) {
  return models.StageData.create(
     item
  );
}

function insertUserProperty(item) {
  return models.UserProperty.upsert(
    item,
      {
      where: {
        UserPropertyID: item.UserPropertyID
      }
  });
}

function insertUserPurchase(item) {
  return models.UserPurchase.create(
    item
  );
}

function insertUserDailyProperty(item) {
  return models.UserDailyProperty.create(
    item
  );
}

function insertTutorial(item) {
  var pKey = parseInt(item.UserPropertyID.toString() + item.TutorialID.toString());
  return models.TutorialData.upsert({
    TutorialDataID: pKey,
    UserPropertyID: item.UserPropertyID,
    TutorialID: item.TutorialID,
    Date: item.Date
  },{
    where: {
      TutorialDataID: pKey
    }
  });
}


module.exports = router;

