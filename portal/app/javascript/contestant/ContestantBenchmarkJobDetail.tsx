import { isuxportal } from "../pb_admin";
import { ApiError, ApiClient } from "../ApiClient";

import React from "react";

import { BenchmarkJobDetail } from "../BenchmarkJobDetail";

import { ErrorMessage } from "../ErrorMessage";
import { ReloadButton } from "../ReloadButton";

export interface Props {
  session: isuxportal.proto.services.common.GetCurrentSessionResponse;
  client: ApiClient;
  id: number | string;
}

export interface State {
  job: isuxportal.proto.resources.IBenchmarkJob | null;
  error: Error | null;
  requesting: boolean;
  timer: number | null;
}

export class ContestantBenchmarkJobDetail extends React.Component<Props, State> {
  constructor(props: Props) {
    super(props);
    this.state = {
      job: null,
      error: null,
      requesting: false,
      timer: null,
    };
  }

  public componentDidMount() {
    this.updateJob();
    this.setState({ timer: window.setInterval(this.updateJob.bind(this), 5000) });
  }

  public componentDidUpdate(prevProps: Props, prevState: State) {
    if (prevProps !== this.props) this.updateJob();
  }

  public componentWillUnmount() {
    if (this.state.timer) window.clearInterval(this.state.timer);
  }

  async updateJob() {
    if (this.state.requesting) return;
    try {
      this.setState({ requesting: true });
      const resp = await this.props.client.getBenchmarkJob(this.props.id);
      this.setState({ job: resp.job!, requesting: false, error: null });
    } catch (error) {
      this.setState({ error, requesting: false });
    }
  }

  public render() {
    return (
      <>
        <header>
          <div className="level">
            <div className="level-left">
              <h1 className="title is-1">Job #{this.props.id}</h1>
            </div>
            <div className="level-right">
              <ReloadButton requesting={this.state.requesting} onClick={this.updateJob.bind(this)} />
            </div>
          </div>
        </header>
        <main>
          {this.renderError()}
          {this.renderJob()}
        </main>
      </>
    );
  }

  public renderError() {
    if (!this.state.error) return;
    return <ErrorMessage error={this.state.error} />;
  }

  renderJob() {
    if (!this.state.job) return <p>Loading...</p>;
    return (
      <>
        <BenchmarkJobDetail job={this.state.job} admin={false} />
      </>
    );
  }
}
