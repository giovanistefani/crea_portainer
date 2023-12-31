import angular from 'angular';

import { EdgeGroupFormController } from './groupFormController';

angular.module('portainer.edge').component('edgeGroupForm', {
  templateUrl: './groupForm.html',
  controller: EdgeGroupFormController,
  bindings: {
    model: '<',
    formActionLabel: '@',
    formAction: '<',
    actionInProgress: '<',
    loaded: '<',
    pageType: '@',
  },
});
