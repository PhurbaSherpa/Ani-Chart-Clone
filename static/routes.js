import React from "react";
import { Route, Switch } from "react-router-dom";
import { AllAnimes, Home, SeasonList } from "./components";

export default function Routes(props) {
  const { animes } = props;
  return (
    <Switch>
      <Route
        path="/Winter"
        render={navProps => <SeasonList {...navProps} animes={animes} />}
      />{" "}
      <Route
        path="/Spring"
        render={navProps => <SeasonList {...navProps} animes={animes} />}
      />{" "}
      <Route
        path="/Summer"
        render={navProps => <SeasonList {...navProps} animes={animes} />}
      />{" "}
      <Route
        path="/Fall"
        render={navProps => <SeasonList {...navProps} animes={animes} />}
      />
      <Route
        path="/animes"
        render={navProps => <AllAnimes {...navProps} animes={animes} />}
      />
      <Route
        path="/"
        render={navProps => <Home {...navProps} animes={animes} />}
      />
    </Switch>
  );
}
