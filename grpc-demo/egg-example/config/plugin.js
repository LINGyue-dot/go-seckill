'use strict';

/** @type Egg.EggPlugin */
module.exports = {
  // had enabled by egg
  // static: {
  //   enable: true,
  // }
    grpc:{
        enable:true,
        package:'egg-grpc'
    }
};
