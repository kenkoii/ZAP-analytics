var express = require('express');
var router = express.Router();
var models = require('../models/index');
var sequelize = models.sequelize;
/* GET home page. */
router.get('/general', function(req, res, next) {
  res.render('general');
});

router.post('/general', function(req, res, next) {
  sequelize.query("SELECT\r\n    s.tag AS \"DATE\",\r\n    count(DISTINCT activity.\"userId\") AS \"DAU\",\r\n    count(DISTINCT activity.\"userId\") - count(DISTINCT installs.\"userId\") AS \"RDAU\",\r\n    count(DISTINCT installs.\"userId\") AS \"Installs\",\r\n    count(DISTINCT one.\"userId\") AS \"1DR\",\r\n    count(DISTINCT three.\"userId\") AS \"3DR\",\r\n    count(DISTINCT seven.\"userId\") AS \"7DR\",\r\n    count(DISTINCT purchase.\"userId\") AS \"Purchase User\",\r\n    COALESCE(NULLIF(sum(DISTINCT udp.price),0), 0) AS \"Purchase Amount\",\r\n    count(DISTINCT purchase.\"userId\") \/ COALESCE(NULLIF(count(DISTINCT activity.\"userId\"),0), 1) AS \"Purchase Rate\",\r\n    COALESCE(NULLIF(max(udp.price),0), 0) AS \"Max Purchase Amount\",\r\n    CAST(COALESCE(NULLIF(sum(DISTINCT udp.price),0), 0) \/ COALESCE(NULLIF(count(DISTINCT purchase.\"userId\"),0), 1) AS DECIMAL(10,2)) AS \"ARPU\",\r\n    CAST(COALESCE(NULLIF(sum(DISTINCT udp.price),0), 0) \/ COALESCE(NULLIF(count(DISTINCT activity.\"userId\"),0), 1) AS DECIMAL(10,2)) AS \"ARPPU\"\r\nFROM (\r\n   SELECT generate_series(?, ?, \'1 day\'::interval)::DATE AS tag\r\n   ) s\r\nLEFT JOIN \"UserDailyProperties\" AS activity ON\r\n  activity.\"LoginDate\"::DATE = s.tag\r\nLEFT JOIN \"UserProperties\" AS installs ON\r\n  installs.\"DownloadDate\"::DATE = s.tag\r\nLEFT JOIN \"UserDailyProperties\" AS one ON\r\n  activity.\"userId\" = one.\"userId\"\r\n  AND activity.\"LoginDate\" = one.\"LoginDate\" + INTERVAL \'1 day\'\r\nLEFT JOIN \"UserDailyProperties\" AS three ON\r\n  activity.\"userId\" = three.\"userId\"\r\n  AND activity.\"LoginDate\" = three.\"LoginDate\" + INTERVAL \'3 day\'\r\nLEFT JOIN \"UserDailyProperties\" AS seven ON\r\n  activity.\"userId\" = seven.\"userId\"\r\n  AND activity.\"LoginDate\" = seven.\"LoginDate\" + INTERVAL \'7 day\'\r\nLEFT JOIN \"UserPurchases\" AS purchase ON\r\n  purchase.\"date\"::DATE = s.tag\r\nLEFT JOIN (\r\n    SELECT\r\n        \"date\", \r\n        SUM(\"price\") AS price\r\n    FROM \"UserPurchases\"\r\n    GROUP BY \"userId\", \"date\"\r\n) AS udp ON\r\n   udp.\"date\"::DATE = s.tag\r\nGROUP BY \"DATE\"", { replacements: [req.body.from, req.body.to],type: sequelize.QueryTypes.SELECT})
  .then(function(results) {
    console.log(results);
    res.locals.results = results;
    res.locals.render = true;
    res.render('general');
  })
});




module.exports = router;

