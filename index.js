var path = require('path');
var os = require('os');
var fs = require('fs');
var exec = require('child_process').exec;

// detect the current OS
var os = os.type().toLowerCase();
var dirname = path.resolve(__dirname);

const BIN_FOLDER = dirname + '/bin/'+os+'-amd64';
const GIT_FOLDER = dirname + '/git/bin';
const GIT_CORE_FOLDER = dirname + '/git/lib/git-core';
const DEFAULT_BRANCH = 'master';
const GO_PATH = '/tmp/gopath';

process.env.PATH = process.env.PATH + ':' + BIN_FOLDER + ':' + GIT_FOLDER + ':' + GO_PATH + '/bin';
process.env.GOGC = 'off';
process.env.GODEBUG = 'sbrk=1';
process.env.GOROOT = BIN_FOLDER + '/goroot';
process.env.GOPATH = GO_PATH;
process.env.GIT_SSL_NO_VERIFY = '1';
process.env.GIT_EXEC_PATH = GIT_CORE_FOLDER;
process.env.GIT_TEMPLATE_DIR = dirname + '/git/templates';
process.env.CGO_ENABLED = '0';
process.env.HOME = dirname;

exports.handler = function(event, context) {
  var fail = function(err) {
    if (!(err instanceof Error)) {
      err = new Error(err);
    }
    context.done(null, {
      success: false,
      errors: {
        runner: err.toString()
      }
    });
  };

  var done = function(json) {
    context.done(null, json);
  };

  if (!event.hasOwnProperty('repository')) {
    fail('The repository is missing');
    return;
  }

  // for now only GitHub is supported
  if (event.repository.indexOf('github.com') === -1) {
    fail('Only GitHub is supported');
    return;
  }

  var flags = [];
  if (event.hasOwnProperty("shallow") && event.shallow) {
    flags.push("--shallow");
  }
  if (event.hasOwnProperty("reference") && event.reference.trim() != "") {
    flags.push("--ref " + event.reference);
  }

  var cmd = 'exago-runner ' + flags.join(" ") + ' ' + event.repository;
  child = exec(cmd, {maxBuffer: 1024 * 2000}, function(error, output) {
    cleanup(function() {
      if (error) {
        fail(error);
        return;
      }
      done(JSON.parse(output));
    });
  });

  child.on('error', console.error);
  child.stderr.on('data', console.error);
  child.stdout.on('data', console.log);
};

var cleanup = function(callback) {  
  var cmd = 'rm -fR ' + GO_PATH;
  child = exec(cmd, null, function() {
    callback();
  });
};