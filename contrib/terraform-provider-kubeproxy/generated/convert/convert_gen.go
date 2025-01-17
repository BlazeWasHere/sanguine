package convert

import (
	diagnostics_context "context"
	schema_context "context"
	schema_fmt "fmt"
	schema_reflect "reflect"
	schema_sort "sort"

	diagnostics_cty "github.com/hashicorp/go-cty/cty"
	schema_cty "github.com/hashicorp/go-cty/cty"
	diagnostics_tfprotov5 "github.com/hashicorp/terraform-plugin-go/tfprotov5"
	schema_tfprotov5 "github.com/hashicorp/terraform-plugin-go/tfprotov5"
	diagnostics_tftypes "github.com/hashicorp/terraform-plugin-go/tftypes"
	schema_tftypes "github.com/hashicorp/terraform-plugin-go/tftypes"
	diagnostics_diag "github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	schema_configschema "github.com/synapsecns/sanguine/contrib/terraform-provider-kubeproxy/generated/configschema"
	diagnostics_logging "github.com/synapsecns/sanguine/contrib/terraform-provider-kubeproxy/generated/logging"
	schema_logging "github.com/synapsecns/sanguine/contrib/terraform-provider-kubeproxy/generated/logging"
)

func AppendProtoDiag(ctx diagnostics_context.Context, diags []*diagnostics_tfprotov5.Diagnostic, d interface{}) []*diagnostics_tfprotov5.Diagnostic {
	switch d := d.(type) {
	case diagnostics_cty.PathError:
		ap := PathToAttributePath(d.Path)
		diagnostic := &diagnostics_tfprotov5.Diagnostic{
			Severity:  diagnostics_tfprotov5.DiagnosticSeverityError,
			Summary:   d.Error(),
			Attribute: ap,
		}

		if diagnostic.Summary == "" {
			diagnostics_logging.HelperSchemaWarn(ctx, "detected empty error string for diagnostic in AppendProtoDiag for cty.PathError type")
			diagnostic.Summary = "Empty Error String"
			diagnostic.Detail = "This is always a bug in the provider and should be reported to the provider developers."
		}

		diags = append(diags, diagnostic)
	case diagnostics_diag.Diagnostics:
		diags = append(diags, DiagsToProto(d)...)
	case error:
		if d == nil {
			diagnostics_logging.HelperSchemaDebug(ctx, "skipping diagnostic for nil error in AppendProtoDiag")
			return diags
		}

		diagnostic := &diagnostics_tfprotov5.Diagnostic{
			Severity: diagnostics_tfprotov5.DiagnosticSeverityError,
			Summary:  d.Error(),
		}

		if diagnostic.Summary == "" {
			diagnostics_logging.HelperSchemaWarn(ctx, "detected empty error string for diagnostic in AppendProtoDiag for error type")
			diagnostic.Summary = "Error Missing Message"
			diagnostic.Detail = "This is always a bug in the provider and should be reported to the provider developers."
		}

		diags = append(diags, diagnostic)
	case string:
		if d == "" {
			diagnostics_logging.HelperSchemaDebug(ctx, "skipping diagnostic for empty string in AppendProtoDiag")
			return diags
		}

		diags = append(diags, &diagnostics_tfprotov5.Diagnostic{
			Severity: diagnostics_tfprotov5.DiagnosticSeverityWarning,
			Summary:  d,
		})
	case *diagnostics_tfprotov5.Diagnostic:
		diags = append(diags, d)
	case []*diagnostics_tfprotov5.Diagnostic:
		diags = append(diags, d...)
	}
	return diags
}

func ProtoToDiags(ds []*diagnostics_tfprotov5.Diagnostic) diagnostics_diag.Diagnostics {
	var diags diagnostics_diag.Diagnostics
	for _, d := range ds {
		var severity diagnostics_diag.Severity

		switch d.Severity {
		case diagnostics_tfprotov5.DiagnosticSeverityError:
			severity = diagnostics_diag.Error
		case diagnostics_tfprotov5.DiagnosticSeverityWarning:
			severity = diagnostics_diag.Warning
		}

		diags = append(diags, diagnostics_diag.Diagnostic{
			Severity:      severity,
			Summary:       d.Summary,
			Detail:        d.Detail,
			AttributePath: AttributePathToPath(d.Attribute),
		})
	}

	return diags
}

func DiagsToProto(diags diagnostics_diag.Diagnostics) []*diagnostics_tfprotov5.Diagnostic {
	var ds []*diagnostics_tfprotov5.Diagnostic
	for _, d := range diags {
		protoDiag := &diagnostics_tfprotov5.Diagnostic{
			Severity:  diagnostics_tfprotov5.DiagnosticSeverityError,
			Summary:   d.Summary,
			Detail:    d.Detail,
			Attribute: PathToAttributePath(d.AttributePath),
		}
		if d.Severity == diagnostics_diag.Warning {
			protoDiag.Severity = diagnostics_tfprotov5.DiagnosticSeverityWarning
		}
		if d.Summary == "" {
			protoDiag.Summary = "Empty Summary: This is always a bug in the provider and should be reported to the provider developers."
		}
		ds = append(ds, protoDiag)
	}
	return ds
}

func AttributePathToPath(ap *diagnostics_tftypes.AttributePath) diagnostics_cty.Path {
	var p diagnostics_cty.Path
	if ap == nil {
		return p
	}
	for _, step := range ap.Steps() {
		switch step := step.(type) {
		case diagnostics_tftypes.AttributeName:
			p = p.GetAttr(string(step))
		case diagnostics_tftypes.ElementKeyString:
			p = p.Index(diagnostics_cty.StringVal(string(step)))
		case diagnostics_tftypes.ElementKeyInt:
			p = p.Index(diagnostics_cty.NumberIntVal(int64(step)))
		}
	}
	return p
}

func PathToAttributePath(p diagnostics_cty.Path) *diagnostics_tftypes.AttributePath {
	if p == nil || len(p) < 1 {
		return nil
	}
	ap := diagnostics_tftypes.NewAttributePath()
	for _, step := range p {
		switch selector := step.(type) {
		case diagnostics_cty.GetAttrStep:
			ap = ap.WithAttributeName(selector.Name)

		case diagnostics_cty.IndexStep:
			key := selector.Key
			switch key.Type() {
			case diagnostics_cty.String:
				ap = ap.WithElementKeyString(key.AsString())
			case diagnostics_cty.Number:
				v, _ := key.AsBigFloat().Int64()
				ap = ap.WithElementKeyInt(int(v))
			default:

				return ap
			}
		}
	}
	return ap
}

func tftypeFromCtyType(in schema_cty.Type) (schema_tftypes.Type, error) {
	switch {
	case in.Equals(schema_cty.String):
		return schema_tftypes.String, nil
	case in.Equals(schema_cty.Number):
		return schema_tftypes.Number, nil
	case in.Equals(schema_cty.Bool):
		return schema_tftypes.Bool, nil
	case in.Equals(schema_cty.DynamicPseudoType):
		return schema_tftypes.DynamicPseudoType, nil
	case in.IsSetType():
		elemType, err := tftypeFromCtyType(in.ElementType())
		if err != nil {
			return nil, err
		}
		return schema_tftypes.Set{
			ElementType: elemType,
		}, nil
	case in.IsListType():
		elemType, err := tftypeFromCtyType(in.ElementType())
		if err != nil {
			return nil, err
		}
		return schema_tftypes.List{
			ElementType: elemType,
		}, nil
	case in.IsTupleType():
		elemTypes := make([]schema_tftypes.Type, 0, in.Length())
		for _, typ := range in.TupleElementTypes() {
			elemType, err := tftypeFromCtyType(typ)
			if err != nil {
				return nil, err
			}
			elemTypes = append(elemTypes, elemType)
		}
		return schema_tftypes.Tuple{
			ElementTypes: elemTypes,
		}, nil
	case in.IsMapType():
		elemType, err := tftypeFromCtyType(in.ElementType())
		if err != nil {
			return nil, err
		}
		return schema_tftypes.Map{
			ElementType: elemType,
		}, nil
	case in.IsObjectType():
		attrTypes := make(map[string]schema_tftypes.Type)
		for key, typ := range in.AttributeTypes() {
			attrType, err := tftypeFromCtyType(typ)
			if err != nil {
				return nil, err
			}
			attrTypes[key] = attrType
		}
		return schema_tftypes.Object{
			AttributeTypes: attrTypes,
		}, nil
	}
	return nil, schema_fmt.Errorf("unknown cty type %s", in.GoString())
}

func ctyTypeFromTFType(in schema_tftypes.Type) (schema_cty.Type, error) {
	switch {
	case in.Is(schema_tftypes.String):
		return schema_cty.String, nil
	case in.Is(schema_tftypes.Bool):
		return schema_cty.Bool, nil
	case in.Is(schema_tftypes.Number):
		return schema_cty.Number, nil
	case in.Is(schema_tftypes.DynamicPseudoType):
		return schema_cty.DynamicPseudoType, nil
	case in.Is(schema_tftypes.List{}):
		elemType, err := ctyTypeFromTFType(in.(schema_tftypes.List).ElementType)
		if err != nil {
			return schema_cty.Type{}, err
		}
		return schema_cty.List(elemType), nil
	case in.Is(schema_tftypes.Set{}):
		elemType, err := ctyTypeFromTFType(in.(schema_tftypes.Set).ElementType)
		if err != nil {
			return schema_cty.Type{}, err
		}
		return schema_cty.Set(elemType), nil
	case in.Is(schema_tftypes.Map{}):
		elemType, err := ctyTypeFromTFType(in.(schema_tftypes.Map).ElementType)
		if err != nil {
			return schema_cty.Type{}, err
		}
		return schema_cty.Map(elemType), nil
	case in.Is(schema_tftypes.Tuple{}):
		elemTypes := make([]schema_cty.Type, 0, len(in.(schema_tftypes.Tuple).ElementTypes))
		for _, typ := range in.(schema_tftypes.Tuple).ElementTypes {
			elemType, err := ctyTypeFromTFType(typ)
			if err != nil {
				return schema_cty.Type{}, err
			}
			elemTypes = append(elemTypes, elemType)
		}
		return schema_cty.Tuple(elemTypes), nil
	case in.Is(schema_tftypes.Object{}):
		attrTypes := make(map[string]schema_cty.Type, len(in.(schema_tftypes.Object).AttributeTypes))
		for k, v := range in.(schema_tftypes.Object).AttributeTypes {
			attrType, err := ctyTypeFromTFType(v)
			if err != nil {
				return schema_cty.Type{}, err
			}
			attrTypes[k] = attrType
		}
		return schema_cty.Object(attrTypes), nil
	}
	return schema_cty.Type{}, schema_fmt.Errorf("unknown tftypes.Type %s", in)
}

func ConfigSchemaToProto(ctx schema_context.Context, b *schema_configschema.Block) *schema_tfprotov5.SchemaBlock {
	block := &schema_tfprotov5.SchemaBlock{
		Description:     b.Description,
		DescriptionKind: protoStringKind(ctx, b.DescriptionKind),
		Deprecated:      b.Deprecated,
	}

	for _, name := range sortedKeys(b.Attributes) {
		a := b.Attributes[name]

		attr := &schema_tfprotov5.SchemaAttribute{
			Name:            name,
			Description:     a.Description,
			DescriptionKind: protoStringKind(ctx, a.DescriptionKind),
			Optional:        a.Optional,
			Computed:        a.Computed,
			Required:        a.Required,
			Sensitive:       a.Sensitive,
			Deprecated:      a.Deprecated,
		}

		var err error
		attr.Type, err = tftypeFromCtyType(a.Type)
		if err != nil {
			panic(err)
		}

		block.Attributes = append(block.Attributes, attr)
	}

	for _, name := range sortedKeys(b.BlockTypes) {
		b := b.BlockTypes[name]
		block.BlockTypes = append(block.BlockTypes, protoSchemaNestedBlock(ctx, name, b))
	}

	return block
}

func protoStringKind(ctx schema_context.Context, k schema_configschema.StringKind) schema_tfprotov5.StringKind {
	switch k {
	default:
		schema_logging.HelperSchemaTrace(ctx, schema_fmt.Sprintf("Unexpected configschema.StringKind: %d", k))
		return schema_tfprotov5.StringKindPlain
	case schema_configschema.StringPlain:
		return schema_tfprotov5.StringKindPlain
	case schema_configschema.StringMarkdown:
		return schema_tfprotov5.StringKindMarkdown
	}
}

func protoSchemaNestedBlock(ctx schema_context.Context, name string, b *schema_configschema.NestedBlock) *schema_tfprotov5.SchemaNestedBlock {
	var nesting schema_tfprotov5.SchemaNestedBlockNestingMode
	switch b.Nesting {
	case schema_configschema.NestingSingle:
		nesting = schema_tfprotov5.SchemaNestedBlockNestingModeSingle
	case schema_configschema.NestingGroup:
		nesting = schema_tfprotov5.SchemaNestedBlockNestingModeGroup
	case schema_configschema.NestingList:
		nesting = schema_tfprotov5.SchemaNestedBlockNestingModeList
	case schema_configschema.NestingSet:
		nesting = schema_tfprotov5.SchemaNestedBlockNestingModeSet
	case schema_configschema.NestingMap:
		nesting = schema_tfprotov5.SchemaNestedBlockNestingModeMap
	default:
		nesting = schema_tfprotov5.SchemaNestedBlockNestingModeInvalid
	}
	return &schema_tfprotov5.SchemaNestedBlock{
		TypeName: name,
		Block:    ConfigSchemaToProto(ctx, &b.Block),
		Nesting:  nesting,
		MinItems: int64(b.MinItems),
		MaxItems: int64(b.MaxItems),
	}
}

func ProtoToConfigSchema(ctx schema_context.Context, b *schema_tfprotov5.SchemaBlock) *schema_configschema.Block {
	block := &schema_configschema.Block{
		Attributes: make(map[string]*schema_configschema.Attribute),
		BlockTypes: make(map[string]*schema_configschema.NestedBlock),

		Description:     b.Description,
		DescriptionKind: schemaStringKind(ctx, b.DescriptionKind),
		Deprecated:      b.Deprecated,
	}

	for _, a := range b.Attributes {
		attr := &schema_configschema.Attribute{
			Description:     a.Description,
			DescriptionKind: schemaStringKind(ctx, a.DescriptionKind),
			Required:        a.Required,
			Optional:        a.Optional,
			Computed:        a.Computed,
			Sensitive:       a.Sensitive,
			Deprecated:      a.Deprecated,
		}

		var err error
		attr.Type, err = ctyTypeFromTFType(a.Type)
		if err != nil {
			panic(err)
		}

		block.Attributes[a.Name] = attr
	}

	for _, b := range b.BlockTypes {
		block.BlockTypes[b.TypeName] = schemaNestedBlock(ctx, b)
	}

	return block
}

func schemaStringKind(ctx schema_context.Context, k schema_tfprotov5.StringKind) schema_configschema.StringKind {
	switch k {
	default:
		schema_logging.HelperSchemaTrace(ctx, schema_fmt.Sprintf("Unexpected tfprotov5.StringKind: %d", k))
		return schema_configschema.StringPlain
	case schema_tfprotov5.StringKindPlain:
		return schema_configschema.StringPlain
	case schema_tfprotov5.StringKindMarkdown:
		return schema_configschema.StringMarkdown
	}
}

func schemaNestedBlock(ctx schema_context.Context, b *schema_tfprotov5.SchemaNestedBlock) *schema_configschema.NestedBlock {
	var nesting schema_configschema.NestingMode
	switch b.Nesting {
	case schema_tfprotov5.SchemaNestedBlockNestingModeSingle:
		nesting = schema_configschema.NestingSingle
	case schema_tfprotov5.SchemaNestedBlockNestingModeGroup:
		nesting = schema_configschema.NestingGroup
	case schema_tfprotov5.SchemaNestedBlockNestingModeList:
		nesting = schema_configschema.NestingList
	case schema_tfprotov5.SchemaNestedBlockNestingModeMap:
		nesting = schema_configschema.NestingMap
	case schema_tfprotov5.SchemaNestedBlockNestingModeSet:
		nesting = schema_configschema.NestingSet
	default:

	}

	nb := &schema_configschema.NestedBlock{
		Nesting:  nesting,
		MinItems: int(b.MinItems),
		MaxItems: int(b.MaxItems),
	}

	nested := ProtoToConfigSchema(ctx, b.Block)
	nb.Block = *nested
	return nb
}

func sortedKeys(m interface{}) []string {
	v := schema_reflect.ValueOf(m)
	keys := make([]string, v.Len())

	mapKeys := v.MapKeys()
	for i, k := range mapKeys {
		keys[i] = k.Interface().(string)
	}

	schema_sort.Strings(keys)
	return keys
}
