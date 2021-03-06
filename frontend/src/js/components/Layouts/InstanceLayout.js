import React from 'react';
import API from '../../api/API';
import { getInstanceStatus } from '../../constants/helpers';
import { applicationsStore } from '../../stores/Stores';
import Loader from '../Common/Loader';
import SectionHeader from '../Common/SectionHeader';
import Details from '../Instances/Details';

export default function InstanceLayout(props) {
  const {appID, groupID, instanceID} = props.match.params;
  const [application, setApplication] = React.useState(
    applicationsStore
      .getCachedApplication(appID)
  );
  const [group, setGroup] = React.useState(getGroupFromApplication(application));
  const [instance, setInstance] = React.useState(null);

  function onChange() {
    API.getInstance(appID, groupID, instanceID)
      .then((instance) => {
        instance.statusInfo = getInstanceStatus(instance.application.status,
          instance.application.version, instance.application.error_code);
        setInstance(instance);
      });
    const apps = applicationsStore.getCachedApplications() || [];
    const app = apps.find(({id}) => id === appID);
    if (app !== application) {
      setApplication(app);
      setGroup(getGroupFromApplication(app));
    }
  }

  function getGroupFromApplication(app) {
    return app ? app.groups.find(({id}) => id === groupID) : null;
  }

  React.useEffect(() => {
    applicationsStore.addChangeListener(onChange);

    applicationsStore.getApplication(appID);

    return function cleanup() {
      applicationsStore.removeChangeListener(onChange);
    };
  },
  []);

  const applicationName = application ? application.name : '…';
  const groupName = group ? group.name : '…';
  const instanceName = (instance && instance.alias) || instanceID;

  const searchParams = new URLSearchParams(window.location.search).toString();
  return (
    <React.Fragment>
      <SectionHeader
        title={instanceName}
        breadcrumbs={[
          {
            path: '/apps',
            label: 'Applications'
          },
          {
            path: `/apps/${appID}`,
            label: applicationName
          },
          {
            path: `/apps/${appID}/groups/${groupID}`,
            label: groupName
          },
          {
            path: `/apps/${appID}/groups/${groupID}/instances?${searchParams}`,
            label: 'Instances'
          },
        ]}
      />
      { !instance ? <Loader />
        :
      <Details
        application={application}
        group={group}
        instance={instance}
        onInstanceUpdated={() => onChange()}
      />
      }
    </React.Fragment>
  );
}
