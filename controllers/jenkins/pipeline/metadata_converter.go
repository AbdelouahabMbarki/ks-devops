package pipeline

import (
	"github.com/jenkins-zh/jenkins-client/pkg/job"
	"kubesphere.io/devops/pkg/models/pipeline"
)

func convertPipeline(jobPipeline *job.Pipeline) *pipeline.Metadata {
	return &pipeline.Metadata{
		WeatherScore:                   jobPipeline.WeatherScore,
		EstimatedDurationInMillis:      jobPipeline.EstimatedDurationInMillis,
		Parameters:                     jobPipeline.Parameters,
		Name:                           jobPipeline.Name,
		Disabled:                       jobPipeline.Disabled,
		NumberOfPipelines:              jobPipeline.NumberOfPipelines,
		NumberOfFolders:                jobPipeline.NumberOfFolders,
		PipelineFolderNames:            jobPipeline.PipelineFolderNames,
		TotalNumberOfBranches:          jobPipeline.TotalNumberOfBranches,
		TotalNumberOfPullRequests:      jobPipeline.TotalNumberOfPullRequests,
		NumberOfFailingBranches:        jobPipeline.NumberOfFailingBranches,
		NumberOfFailingPullRequests:    jobPipeline.NumberOfFailingPullRequests,
		NumberOfSuccessfulBranches:     jobPipeline.NumberOfSuccessfulBranches,
		NumberOfSuccessfulPullRequests: jobPipeline.NumberOfSuccessfulPullRequests,
		BranchNames:                    jobPipeline.BranchNames,
		SCMSource:                      jobPipeline.SCMSource,
		ScriptPath:                     jobPipeline.ScriptPath,
	}
}

func convertLatestRun(jobLatestRun *job.PipelineRunSummary) *pipeline.LatestRun {
	if jobLatestRun == nil {
		return nil
	}
	return &pipeline.LatestRun{
		ID:        jobLatestRun.ID,
		Name:      jobLatestRun.Name,
		Pipeline:  jobLatestRun.Pipeline,
		Result:    jobLatestRun.Result,
		State:     jobLatestRun.State,
		StartTime: jobLatestRun.StartTime,
		EndTime:   jobLatestRun.EndTime,
		Causes:    convertCauses(jobLatestRun.Causes),
	}
}

func convertCauses(jobCauses []job.Cause) []pipeline.Cause {
	if jobCauses == nil {
		return nil
	}
	causes := []pipeline.Cause{}
	for _, jobCause := range jobCauses {
		causes = append(causes, pipeline.Cause{
			ShortDescription: jobCause.GetShortDescription(),
		})
	}
	return causes
}

func convertBranches(jobBranches []job.PipelineBranch) []pipeline.Branch {
	branches := make([]pipeline.Branch, 0, len(jobBranches))
	for _, jobBranch := range jobBranches {
		branches = append(branches, pipeline.Branch{
			Name:         jobBranch.Name,
			WeatherScore: jobBranch.WeatherScore,
			Branch:       jobBranch.Branch,
			PullRequest:  jobBranch.PullRequest,
			LatestRun:    convertLatestRun(jobBranch.LatestRun),
		})
	}
	return branches
}