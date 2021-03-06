// Code generated by entc, DO NOT EDIT.

package ent

import (
	"github.com/NpoolPlatform/innovation-minning/pkg/db/ent/capital"
	"github.com/NpoolPlatform/innovation-minning/pkg/db/ent/launchtime"
	"github.com/NpoolPlatform/innovation-minning/pkg/db/ent/project"
	"github.com/NpoolPlatform/innovation-minning/pkg/db/ent/schema"
	"github.com/NpoolPlatform/innovation-minning/pkg/db/ent/team"
	"github.com/NpoolPlatform/innovation-minning/pkg/db/ent/techniqueanalysis"
	"github.com/NpoolPlatform/innovation-minning/pkg/db/ent/trendanalysis"
	"github.com/NpoolPlatform/innovation-minning/pkg/db/ent/user"
	"github.com/google/uuid"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	capitalFields := schema.Capital{}.Fields()
	_ = capitalFields
	// capitalDescCreateAt is the schema descriptor for create_at field.
	capitalDescCreateAt := capitalFields[4].Descriptor()
	// capital.DefaultCreateAt holds the default value on creation for the create_at field.
	capital.DefaultCreateAt = capitalDescCreateAt.Default.(func() uint32)
	// capitalDescUpdateAt is the schema descriptor for update_at field.
	capitalDescUpdateAt := capitalFields[5].Descriptor()
	// capital.DefaultUpdateAt holds the default value on creation for the update_at field.
	capital.DefaultUpdateAt = capitalDescUpdateAt.Default.(func() uint32)
	// capital.UpdateDefaultUpdateAt holds the default value on update for the update_at field.
	capital.UpdateDefaultUpdateAt = capitalDescUpdateAt.UpdateDefault.(func() uint32)
	// capitalDescDeleteAt is the schema descriptor for delete_at field.
	capitalDescDeleteAt := capitalFields[6].Descriptor()
	// capital.DefaultDeleteAt holds the default value on creation for the delete_at field.
	capital.DefaultDeleteAt = capitalDescDeleteAt.Default.(func() uint32)
	// capitalDescID is the schema descriptor for id field.
	capitalDescID := capitalFields[0].Descriptor()
	// capital.DefaultID holds the default value on creation for the id field.
	capital.DefaultID = capitalDescID.Default.(func() uuid.UUID)
	launchtimeFields := schema.LaunchTime{}.Fields()
	_ = launchtimeFields
	// launchtimeDescCreateAt is the schema descriptor for create_at field.
	launchtimeDescCreateAt := launchtimeFields[8].Descriptor()
	// launchtime.DefaultCreateAt holds the default value on creation for the create_at field.
	launchtime.DefaultCreateAt = launchtimeDescCreateAt.Default.(func() uint32)
	// launchtimeDescUpdateAt is the schema descriptor for update_at field.
	launchtimeDescUpdateAt := launchtimeFields[9].Descriptor()
	// launchtime.DefaultUpdateAt holds the default value on creation for the update_at field.
	launchtime.DefaultUpdateAt = launchtimeDescUpdateAt.Default.(func() uint32)
	// launchtime.UpdateDefaultUpdateAt holds the default value on update for the update_at field.
	launchtime.UpdateDefaultUpdateAt = launchtimeDescUpdateAt.UpdateDefault.(func() uint32)
	// launchtimeDescDeleteAt is the schema descriptor for delete_at field.
	launchtimeDescDeleteAt := launchtimeFields[10].Descriptor()
	// launchtime.DefaultDeleteAt holds the default value on creation for the delete_at field.
	launchtime.DefaultDeleteAt = launchtimeDescDeleteAt.Default.(func() uint32)
	// launchtimeDescID is the schema descriptor for id field.
	launchtimeDescID := launchtimeFields[0].Descriptor()
	// launchtime.DefaultID holds the default value on creation for the id field.
	launchtime.DefaultID = launchtimeDescID.Default.(func() uuid.UUID)
	projectFields := schema.Project{}.Fields()
	_ = projectFields
	// projectDescCreateAt is the schema descriptor for create_at field.
	projectDescCreateAt := projectFields[7].Descriptor()
	// project.DefaultCreateAt holds the default value on creation for the create_at field.
	project.DefaultCreateAt = projectDescCreateAt.Default.(func() uint32)
	// projectDescUpdateAt is the schema descriptor for update_at field.
	projectDescUpdateAt := projectFields[8].Descriptor()
	// project.DefaultUpdateAt holds the default value on creation for the update_at field.
	project.DefaultUpdateAt = projectDescUpdateAt.Default.(func() uint32)
	// project.UpdateDefaultUpdateAt holds the default value on update for the update_at field.
	project.UpdateDefaultUpdateAt = projectDescUpdateAt.UpdateDefault.(func() uint32)
	// projectDescDeleteAt is the schema descriptor for delete_at field.
	projectDescDeleteAt := projectFields[9].Descriptor()
	// project.DefaultDeleteAt holds the default value on creation for the delete_at field.
	project.DefaultDeleteAt = projectDescDeleteAt.Default.(func() uint32)
	// projectDescID is the schema descriptor for id field.
	projectDescID := projectFields[0].Descriptor()
	// project.DefaultID holds the default value on creation for the id field.
	project.DefaultID = projectDescID.Default.(func() uuid.UUID)
	teamFields := schema.Team{}.Fields()
	_ = teamFields
	// teamDescCreateAt is the schema descriptor for create_at field.
	teamDescCreateAt := teamFields[6].Descriptor()
	// team.DefaultCreateAt holds the default value on creation for the create_at field.
	team.DefaultCreateAt = teamDescCreateAt.Default.(func() uint32)
	// teamDescUpdateAt is the schema descriptor for update_at field.
	teamDescUpdateAt := teamFields[7].Descriptor()
	// team.DefaultUpdateAt holds the default value on creation for the update_at field.
	team.DefaultUpdateAt = teamDescUpdateAt.Default.(func() uint32)
	// team.UpdateDefaultUpdateAt holds the default value on update for the update_at field.
	team.UpdateDefaultUpdateAt = teamDescUpdateAt.UpdateDefault.(func() uint32)
	// teamDescDeleteAt is the schema descriptor for delete_at field.
	teamDescDeleteAt := teamFields[8].Descriptor()
	// team.DefaultDeleteAt holds the default value on creation for the delete_at field.
	team.DefaultDeleteAt = teamDescDeleteAt.Default.(func() uint32)
	// teamDescID is the schema descriptor for id field.
	teamDescID := teamFields[0].Descriptor()
	// team.DefaultID holds the default value on creation for the id field.
	team.DefaultID = teamDescID.Default.(func() uuid.UUID)
	techniqueanalysisFields := schema.TechniqueAnalysis{}.Fields()
	_ = techniqueanalysisFields
	// techniqueanalysisDescCreateAt is the schema descriptor for create_at field.
	techniqueanalysisDescCreateAt := techniqueanalysisFields[6].Descriptor()
	// techniqueanalysis.DefaultCreateAt holds the default value on creation for the create_at field.
	techniqueanalysis.DefaultCreateAt = techniqueanalysisDescCreateAt.Default.(func() uint32)
	// techniqueanalysisDescUpdateAt is the schema descriptor for update_at field.
	techniqueanalysisDescUpdateAt := techniqueanalysisFields[7].Descriptor()
	// techniqueanalysis.DefaultUpdateAt holds the default value on creation for the update_at field.
	techniqueanalysis.DefaultUpdateAt = techniqueanalysisDescUpdateAt.Default.(func() uint32)
	// techniqueanalysis.UpdateDefaultUpdateAt holds the default value on update for the update_at field.
	techniqueanalysis.UpdateDefaultUpdateAt = techniqueanalysisDescUpdateAt.UpdateDefault.(func() uint32)
	// techniqueanalysisDescDeleteAt is the schema descriptor for delete_at field.
	techniqueanalysisDescDeleteAt := techniqueanalysisFields[8].Descriptor()
	// techniqueanalysis.DefaultDeleteAt holds the default value on creation for the delete_at field.
	techniqueanalysis.DefaultDeleteAt = techniqueanalysisDescDeleteAt.Default.(func() uint32)
	// techniqueanalysisDescID is the schema descriptor for id field.
	techniqueanalysisDescID := techniqueanalysisFields[0].Descriptor()
	// techniqueanalysis.DefaultID holds the default value on creation for the id field.
	techniqueanalysis.DefaultID = techniqueanalysisDescID.Default.(func() uuid.UUID)
	trendanalysisFields := schema.TrendAnalysis{}.Fields()
	_ = trendanalysisFields
	// trendanalysisDescCreateAt is the schema descriptor for create_at field.
	trendanalysisDescCreateAt := trendanalysisFields[6].Descriptor()
	// trendanalysis.DefaultCreateAt holds the default value on creation for the create_at field.
	trendanalysis.DefaultCreateAt = trendanalysisDescCreateAt.Default.(func() uint32)
	// trendanalysisDescUpdateAt is the schema descriptor for update_at field.
	trendanalysisDescUpdateAt := trendanalysisFields[7].Descriptor()
	// trendanalysis.DefaultUpdateAt holds the default value on creation for the update_at field.
	trendanalysis.DefaultUpdateAt = trendanalysisDescUpdateAt.Default.(func() uint32)
	// trendanalysis.UpdateDefaultUpdateAt holds the default value on update for the update_at field.
	trendanalysis.UpdateDefaultUpdateAt = trendanalysisDescUpdateAt.UpdateDefault.(func() uint32)
	// trendanalysisDescDeleteAt is the schema descriptor for delete_at field.
	trendanalysisDescDeleteAt := trendanalysisFields[8].Descriptor()
	// trendanalysis.DefaultDeleteAt holds the default value on creation for the delete_at field.
	trendanalysis.DefaultDeleteAt = trendanalysisDescDeleteAt.Default.(func() uint32)
	// trendanalysisDescID is the schema descriptor for id field.
	trendanalysisDescID := trendanalysisFields[0].Descriptor()
	// trendanalysis.DefaultID holds the default value on creation for the id field.
	trendanalysis.DefaultID = trendanalysisDescID.Default.(func() uuid.UUID)
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescCreateAt is the schema descriptor for create_at field.
	userDescCreateAt := userFields[4].Descriptor()
	// user.DefaultCreateAt holds the default value on creation for the create_at field.
	user.DefaultCreateAt = userDescCreateAt.Default.(func() uint32)
	// userDescUpdateAt is the schema descriptor for update_at field.
	userDescUpdateAt := userFields[5].Descriptor()
	// user.DefaultUpdateAt holds the default value on creation for the update_at field.
	user.DefaultUpdateAt = userDescUpdateAt.Default.(func() uint32)
	// user.UpdateDefaultUpdateAt holds the default value on update for the update_at field.
	user.UpdateDefaultUpdateAt = userDescUpdateAt.UpdateDefault.(func() uint32)
	// userDescDeleteAt is the schema descriptor for delete_at field.
	userDescDeleteAt := userFields[6].Descriptor()
	// user.DefaultDeleteAt holds the default value on creation for the delete_at field.
	user.DefaultDeleteAt = userDescDeleteAt.Default.(func() uint32)
	// userDescID is the schema descriptor for id field.
	userDescID := userFields[0].Descriptor()
	// user.DefaultID holds the default value on creation for the id field.
	user.DefaultID = userDescID.Default.(func() uuid.UUID)
}
