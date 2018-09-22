import React from "react";
import PropTypes from "prop-types";
import { BrowserRouter, Route, Switch } from "react-router-dom";
import isEmpty from "lodash/isEmpty";
import NavBar from "../../containers/Navbar";

import Loader from "./Loader";
import StepNumbers from "./StepNumbers";
import DetermineComponentForRoute from "../../containers/DetermineComponentForRoute";
import StepDone from "./StepDone";

const isRootPath = (basePath) => {
  const formattedBasePath = basePath === "/" ? basePath : basePath.replace(/\/$/, "");
  return window.location.pathname === formattedBasePath
}

export default class RouteDecider extends React.Component {
  static propTypes = {
    basePath: PropTypes.string.isRequired,
    isDone: PropTypes.bool.isRequired,
    routes: PropTypes.arrayOf(
      PropTypes.shape({
         id: PropTypes.string,
         description: PropTypes.string,
         phase: PropTypes.string,
      })
    )
  }

  componentDidUpdate(lastProps) {
    const {
      basePath,
      routes,
      getHelmChartMetadata,
    } = this.props
    if (routes !== lastProps.routes && routes.length) {
      for (let i = 0; i < routes.length; i++) {
        if (routes[i].phase.includes("helm")) {
          getHelmChartMetadata();
          break;
        }
      }

      if (isRootPath(basePath)) {
        window.location.replace(`${basePath}${routes[0].id}`);
      }
    }
  }

  componentDidMount() {
    if (isEmpty(this.props.routes)) {
      this.props.getRoutes();
    }
  }

  render() {
    const {
      basePath,
      routes,
      isDone
    } = this.props;
    const isOnRoot = isRootPath(basePath)

    return (
      <div className="u-minHeight--full u-minWidth--full flex-column flex1">
        <BrowserRouter basename={basePath}>
          <div className="flex-column flex1">
            <div className="flex-column flex1 u-overflow--hidden u-position--relative">
              {!routes ?
                <div className="flex1 flex-column justifyContent--center alignItems--center">
                  <Loader size="60" />
                </div>
                :
                <div className="u-minHeight--full u-minWidth--full flex-column flex1">
                  {isOnRoot ? null : <NavBar hideLinks={true} routes={routes} />}
                  {isOnRoot || isDone ? null : <StepNumbers steps={routes} />}
                  <div className="flex-1-auto flex-column u-overflow--auto">
                    <Switch>
                      {routes && routes.map((route) => (
                        <Route
                          exact
                          key={route.id}
                          path={`/${route.id}`}
                          render={() => <DetermineComponentForRoute routes={routes} routeId={route.id} />}
                        />
                      ))}
                      <Route exact path="/" component={() => <div className="flex1 flex-column justifyContent--center alignItems--center"><Loader size="60" /></div> } />
                      <Route exact path="/done" component={() =>  <StepDone />} />
                    </Switch>
                  </div>
                </div>
              }
            </div>
          </div>
        </BrowserRouter>
      </div>
    );
  }
}