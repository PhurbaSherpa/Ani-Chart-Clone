import React from "react";
import { Route, Switch } from "react-router-dom";
import { AllAnimes, Home } from "./components";

export default function Routes() {
  return (
    <Switch>
      <Route path="/animes" component={AllAnimes} />
      <Route path="/" component={Home} />
    </Switch>
  );
}
