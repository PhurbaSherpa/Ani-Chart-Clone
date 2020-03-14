import React from "react";
import { Route, Switch } from "react-router-dom";
import { AllAnimes, Home } from "./components";

export default function Routes(props) {
  const { animes } = props;
  return (
    <Switch>
      <Route
        path="/animes"
        render={navProps => <AllAnimes {...navProps} animes={animes} />}
      />
      <Route path="/" component={Home} />
    </Switch>
  );
}
