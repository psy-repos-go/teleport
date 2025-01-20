/**
 * Teleport
 * Copyright (C) 2024 Gravitational, Inc.
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Affero General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU Affero General Public License for more details.
 *
 * You should have received a copy of the GNU Affero General Public License
 * along with this program.  If not, see <http://www.gnu.org/licenses/>.
 */

import { Route, Switch } from 'teleport/components/Router';
import cfg from 'teleport/config';
import { AwsOidcStatusProvider } from 'teleport/Integrations/status/AwsOidc/useAwsOidcStatus';

import { AwsOidcDashboard } from './AwsOidcDashboard';

export function AwsOidcRoutes() {
  return (
    <AwsOidcStatusProvider>
      <Switch>
        <Route
          key="aws-oidc-resources-list"
          exact
          path={cfg.routes.integrationStatus}
          component={AwsOidcDashboard}
        />
      </Switch>
    </AwsOidcStatusProvider>
  );
}
