var express = require('express');
var router = express.Router();
var models = require('../models/index');
var sequelize = models.sequelize;
/* GET home page. */
router.get('/general', function(req, res, next) {
  sequelize.query("SELECT\r\n    s.tag AS \"DATE\",\r\n    count(DISTINCT activity.\"UserPropertyID\") AS \"DAU\",\r\n    count(DISTINCT activity.\"UserPropertyID\") - count(DISTINCT installs.\"UserPropertyID\") AS \"RDAU\",\r\n    count(DISTINCT installs.\"UserPropertyID\") AS \"Installs\",\r\n    count(DISTINCT one.\"UserPropertyID\") AS \"1DR\",\r\n    count(DISTINCT three.\"UserPropertyID\") AS \"3DR\",\r\n    count(DISTINCT seven.\"UserPropertyID\") AS \"7DR\",\r\n    count(DISTINCT purchase.\"UserPropertyID\") AS \"Purchase User\",\r\n    COALESCE(NULLIF(sum(DISTINCT udp.price),0), 0) AS \"Purchase Amount\",\r\n    count(DISTINCT purchase.\"UserPropertyID\") \/ COALESCE(NULLIF(count(DISTINCT activity.\"UserPropertyID\"),0), 1) AS \"Purchase Rate\",\r\n    COALESCE(NULLIF(max(udp.price),0), 0) AS \"MAX Purchase Amount\",\r\n    COALESCE(NULLIF(sum(DISTINCT udp.price),0), 0) \/ COALESCE(NULLIF(count(DISTINCT purchase.\"UserPropertyID\"),0), 1) AS \"ARPU\",\r\n    COALESCE(NULLIF(sum(DISTINCT udp.price),0), 0) \/ COALESCE(NULLIF(count(DISTINCT activity.\"UserPropertyID\"),0), 1) AS \"ARPPU\"\r\nFROM (\r\n   SELECT generate_series(\'2017-05-08\', \'2017-05-14\', \'1 day\'::interval)::DATE AS tag\r\n   ) s\r\nLEFT JOIN \"UserDailyProperties\" AS activity ON\r\n  activity.\"LoginDate\"::DATE = s.tag\r\nLEFT JOIN \"UserProperties\" AS installs ON\r\n  installs.\"DownloadDate\"::DATE = s.tag\r\nLEFT JOIN \"UserDailyProperties\" AS one ON\r\n  activity.\"UserPropertyID\" = one.\"UserPropertyID\"\r\n  AND activity.\"LoginDate\" = one.\"LoginDate\" + INTERVAL \'1 day\'\r\nLEFT JOIN \"UserDailyProperties\" AS three ON\r\n  activity.\"UserPropertyID\" = three.\"UserPropertyID\"\r\n  AND activity.\"LoginDate\" = three.\"LoginDate\" + INTERVAL \'3 day\'\r\nLEFT JOIN \"UserDailyProperties\" AS seven ON\r\n  activity.\"UserPropertyID\" = seven.\"UserPropertyID\"\r\n  AND activity.\"LoginDate\" = seven.\"LoginDate\" + INTERVAL \'7 day\'\r\nLEFT JOIN \"UserPurchases\" AS purchase ON\r\n  purchase.\"Date\"::DATE = s.tag\r\nLEFT JOIN (\r\n    SELECT\r\n        \"Date\", \r\n        SUM(\"Price\") AS price\r\n    FROM \"UserPurchases\"\r\n    GROUP BY \"UserPropertyID\", \"Date\"\r\n) AS udp ON\r\n   udp.\"Date\"::DATE = s.tag\r\nGROUP BY \"DATE\"", { type: sequelize.QueryTypes.SELECT})
  .then(function(results) {
    console.log(results);
    res.locals.results = results;
    res.render('general');
    // We don't need spread here, since only the results will be returned for select queries
  })
  // res.render('index', { title: 'Express' });
});




module.exports = router;

