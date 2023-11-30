require("dotenv").config({ path: __dirname + '/.env' });
var createError = require('http-errors');
var express = require('express');
var path = require('path');
var cookieParser = require('cookie-parser');
var logger = require('morgan');
const mongoose = require("mongoose");
var cors = require('cors')

const connectDB = require("./config/db");

process.env.TZ = 'Asia/Tehran';

var app = express();

// mongodb configuration
connectDB();

// CronJob Configuration
// const {sendUserGoogleSheetCronJob} = require("./services/cronjob");
// sendUserGoogleSheetCronJob();
// view engine setup
app.set('views', path.join(__dirname, 'views'));
app.set('view engine', 'ejs');

app.use(cors())
app.use(logger('dev'));
app.use(express.json());
app.use(express.urlencoded({ extended: false }));
app.use(cookieParser());
app.use(express.static(path.join(__dirname, 'public')));

var apiRouter = require('./routes/api');
app.use('/api', apiRouter);

// catch 404 and forward to error handler
app.use(function(req, res, next) {
  next(createError(404,"404"));
});

// error handler
app.use(function(err, req, res, next) {
  // set locals, only providing error in development
  res.locals.message = err.message;
  res.locals.error = req.app.get('env') === 'development' ? err : {};

  // render the error page
  res.status(err.status || 500);
  res.render('error');
});

module.exports = app;
