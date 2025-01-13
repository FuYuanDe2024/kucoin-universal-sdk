package com.kucoin.universal.sdk.plugin.generator;

import com.kucoin.universal.sdk.plugin.model.EnumEntry;
import com.kucoin.universal.sdk.plugin.model.Meta;
import com.kucoin.universal.sdk.plugin.model.ModeSwitch;
import com.kucoin.universal.sdk.plugin.service.NameService;
import com.kucoin.universal.sdk.plugin.service.OperationService;
import com.kucoin.universal.sdk.plugin.service.SchemaService;
import com.kucoin.universal.sdk.plugin.service.impl.OperationServiceImpl;
import com.kucoin.universal.sdk.plugin.service.impl.SchemaServiceImpl;
import com.kucoin.universal.sdk.plugin.util.SpecificationUtil;
import io.swagger.v3.oas.models.OpenAPI;
import io.swagger.v3.oas.models.Operation;
import io.swagger.v3.oas.models.media.Schema;
import io.swagger.v3.oas.models.servers.Server;
import lombok.extern.slf4j.Slf4j;
import org.apache.commons.lang3.StringUtils;
import org.openapitools.codegen.*;
import org.openapitools.codegen.languages.AbstractTypeScriptClientCodegen;
import org.openapitools.codegen.model.ModelMap;
import org.openapitools.codegen.model.ModelsMap;
import org.openapitools.codegen.model.OperationMap;
import org.openapitools.codegen.model.OperationsMap;
import org.openapitools.codegen.utils.CamelizeOption;

import java.io.File;
import java.util.*;
import java.util.stream.Collectors;

import static org.openapitools.codegen.utils.StringUtils.camelize;
import static org.openapitools.codegen.utils.StringUtils.underscore;

/**
 * @author isaac.tang
 */
@Slf4j
public class NodeSdkGenerator extends AbstractTypeScriptClientCodegen implements NameService {

    private static class ImportModel {
        public String className;
        public Set<String> component;

        public ImportModel(String className) {
            this.className = className;
            this.component = new HashSet<>();
        }

        public String toImport() {
            if (component.isEmpty()) {
                return String.format("import \"%s\"", className);
            }

            return String.format("import { %s } from \"%s\"", String.join(", ", component), className);
        }
    }


    private SchemaService schemaService;
    private OperationService operationService;
    private ModeSwitch modeSwitch;

    private String service;
    private String subService;
    private Set<String> exports = new LinkedHashSet<>();

    public CodegenType getTag() {
        return CodegenType.OTHER;
    }

    public String getName() {
        return "node-sdk";
    }

    public String getHelp() {
        return "Generates a node-sdk library.";
    }

    public NodeSdkGenerator() {
        super();
        cliOptions.add(ModeSwitch.option);
    }

    @Override
    public void processOpts() {
        super.processOpts();
        modeSwitch = new ModeSwitch(additionalProperties);
        service = openAPI.getInfo().getTitle();
        subService = openAPI.getInfo().getDescription();

        switch (modeSwitch.getMode()) {
            case API: {
                modelTemplateFiles.put("model.mustache", ".ts");
                apiTemplateFiles.put("api.mustache", ".ts");
                supportingFiles.add(new SupportingFile("module.mustache", String.format("./%s/%s/index.ts", service, formatPackage(subService))));
                break;
            }
            case TEST: {
                apiTemplateFiles.put("api_test.mustache", ".ts");
                break;
            }
            case TEST_TEMPLATE: {
                apiTemplateFiles.put("api_test_template.mustache", ".template");
                break;
            }
            case ENTRY: {
                apiTemplateFiles.put("api_entry.mustache", ".ts");
                supportingFiles.add(new SupportingFile("module.mustache", String.format("./%s/index.ts", service)));
                break;
            }
            case WS: {
                modelTemplateFiles.put("model_ws.mustache", ".ts");
                apiTemplateFiles.put("api_ws.mustache", ".ts");
                additionalProperties.put("WS_MODE", "true");
                supportingFiles.add(new SupportingFile("module.mustache", String.format("./%s/%s/index.ts", service, formatPackage(subService))));
                break;
            }
            case WS_TEST: {
                apiTemplateFiles.put("api_ws_test.mustache", ".ts");
                break;
            }
            default:
                throw new RuntimeException("unsupported mode");
        }

        supportingFiles.add(new SupportingFile("version.mustache", "version.ts"));

        templateDir = "node-sdk";

        // override parent properties
        enablePostProcessFile = true;

        inlineSchemaOption.put("SKIP_SCHEMA_REUSE", "true");
    }

    @Override
    public void preprocessOpenAPI(OpenAPI openAPI) {
        super.preprocessOpenAPI(openAPI);

        // parse and update operations and models
        schemaService = new SchemaServiceImpl(openAPI);
        operationService = new OperationServiceImpl(openAPI, this);

        operationService.parseOperation();
        schemaService.parseSchema();
    }

    @Override
    public String formatParamName(String name) {
        return toParamName(name);
    }

    @Override
    public String formatMethodName(String name) {
        return sanitizeName(name);
    }

    @Override
    public String formatService(String name) {
        return camelize(name);
    }

    @Override
    public String formatPackage(String name) {
        return formatService(name).toLowerCase();
    }

    @Override
    public CodegenProperty fromProperty(String name, Schema p, boolean required) {
        CodegenProperty prop = super.fromProperty(name, p, required);
        if (prop.defaultValue != null && prop.defaultValue.equalsIgnoreCase("undefined")) {
            prop.defaultValue = null;
        }

        if (prop.isEnum) {
            List<EnumEntry> enums = new ArrayList<>();

            List<Map<String, Object>> enumList;
            if (prop.openApiType.equalsIgnoreCase("array")) {
                enumList = (List<Map<String, Object>>) prop.mostInnerItems.vendorExtensions.get("x-api-enum");
            } else {
                enumList = (List<Map<String, Object>>) prop.vendorExtensions.get("x-api-enum");
            }


            List<String> names = new ArrayList<>();
            List<String> values = new ArrayList<>();
            List<String> description = new ArrayList<>();

            enumList.forEach(e -> {
                Object enumValueOriginal = e.get("value");

                String enumValueNameGauss;
                if (enumValueOriginal instanceof Integer) {
                    enumValueNameGauss = "_" + e.get("value");
                } else if (enumValueOriginal instanceof String) {
                    enumValueNameGauss = enumValueOriginal.toString();
                } else {
                    throw new IllegalArgumentException("unknown enum value type..." + e.get("value"));
                }

                String enumName = (String) e.get("name");
                if (StringUtils.isEmpty(enumName)) {
                    enumName = enumValueNameGauss;
                }

                enumName = toVarName(enumName).toUpperCase();
                String enumValue = toEnumValue(enumValueOriginal.toString().trim(), typeMapping.get(p.getType()));

                names.add(enumName);
                values.add(enumValueOriginal.toString().trim());
                description.add(e.get("description").toString());

                enums.add(new EnumEntry(enumName, enumValue, enumValueOriginal, (String) e.get("description"), enumValueOriginal instanceof String));
            });

            // update internal enum support
            prop._enum = values;
            prop.allowableValues.put("values", values);
            prop.vendorExtensions.put("x-enum-varnames", names);
            prop.vendorExtensions.put("x-enum-descriptions", description);
            prop.vendorExtensions.put("x-enums", enums);
        }

        return prop;
    }

    public boolean isDataTypeString(String dataType) {
        return "string".equals(dataType);
    }

    @Override
    public String toEnumDefaultValue(String value, String datatype) {
        return datatype + "." + value;
    }

    @Override
    public String getEnumDefaultValue(String defaultValue, String dataType) {
        return defaultValue;
    }

    @Override
    public String toModelName(String name) {
        return formatService(schemaService.getGeneratedModelName(name));
    }

    @Override
    public String toApiName(String name) {
        return camelize(name + "_" + (modeSwitch.isWs() ? "WS" : "API"));
    }

    @Override
    public String toModelFilename(String name) {
        name = schemaService.getGeneratedModelName(name);
        name = "model_" + name;
        name = underscore(name);
        return name;
    }

    @Override
    public String modelFileFolder() {
        switch (modeSwitch.getMode()) {
            case ENTRY:
                return outputFolder + File.separator + "service";
            default:
                return outputFolder + File.separator + service + File.separator + formatPackage(subService);
        }
    }

    @Override
    public String toApiFilename(String name) {
        String apiName = name.replaceAll("-", "_");
        switch (modeSwitch.getMode()) {
            case WS: {
                apiName = "ws_" + underscore(apiName);
                break;
            }
            case API:
            case ENTRY:
            case TEST_TEMPLATE: {
                apiName = "api_" + underscore(apiName);
                break;
            }
            case TEST: {
                apiName = "api_" + underscore(apiName) + ".test";
                break;
            }
            case WS_TEST: {
                apiName = "ws_" + underscore(apiName) + ".test";
                break;
            }
        }

        return apiName;
    }

    @Override
    public String modelFilename(String templateName, String name) {
        String suffix = modelTemplateFiles().get(templateName);
        return modelFileFolder() + File.separator + toModelFilename(name) + suffix;
    }

    @Override
    public String apiFilename(String templateName, String tag) {
        String suffix = apiTemplateFiles().get(templateName);
        if (modeSwitch.isEntry()) {
            String entryType = service + "_api";
            return modelFileFolder() + File.separator + entryType + suffix;
        }
        return modelFileFolder() + File.separator + toApiFilename(tag) + suffix;
    }

    @Override
    public CodegenOperation fromOperation(String path, String httpMethod, Operation operation, List<Server> servers) {
        CodegenOperation result = super.fromOperation(path, httpMethod, operation, servers);
        if (httpMethod.equalsIgnoreCase("patch")) {
            result.httpMethod = (String) operation.getExtensions().get("x-original-method");
        }
        return result;
    }

    private void generateApiImport(Meta meta, boolean req, Map<String, ImportModel> imports) {
        switch (modeSwitch.getMode()) {
            case API:
            case TEST: {
                String suffix = "resp";
                if (req) {
                    suffix = "req";
                }
                String fileName = "./" + toModelFilename(meta.getMethod()) + "_" + suffix;
                String service = formatService(meta.getMethod() + camelize(suffix));

                imports.computeIfAbsent(fileName, ImportModel::new).component.add(service);
                break;
            }
            case WS: {
                String suffix = "event";
                String fileName = "./" + toModelFilename(meta.getMethod()) + "_" + suffix;
                String service1 = formatService(meta.getMethod() + camelize(suffix + "Callback"));
                String service2 = formatService(meta.getMethod() + camelize(suffix + "CallbackWrapper"));

                imports.computeIfAbsent(fileName, ImportModel::new).component.addAll(Arrays.asList(service1, service2));
                break;

            }
            case WS_TEST: {
                String suffix = "event";
                String fileName = "./" + toModelFilename(meta.getMethod()) + "_" + suffix;
                String service = formatService(meta.getMethod() + camelize(suffix));
                imports.computeIfAbsent(fileName, ImportModel::new).component.add(service);
                break;
            }
            case ENTRY: {
                operationService.getServiceMeta().forEach((k, v) -> {
                    if (v.getService().equalsIgnoreCase(meta.getService())) {
                        List<String> services = Arrays.asList(formatService(k + "API"), formatService(k + "APIImpl"));
                        String fileName = String.format("@generate/%s/%s/%s", formatPackage(v.getService()),
                                formatPackage(v.getSubService()), toApiFilename(formatMethodName(k)));
                        imports.computeIfAbsent(fileName, ImportModel::new).component.addAll(services);
                    }
                });
                break;
            }
            default: {
                throw new RuntimeException("unsupported mode");
            }
        }
    }

    private void generateApiExport(Meta meta, boolean req, Set<String> export) {
        switch (modeSwitch.getMode()) {
            case API: {
                String suffix = "resp";
                if (req) {
                    suffix = "req";
                }
                export.add(String.format("export * from \"./%s\"", toModelFilename(meta.getMethod()) + "_" + suffix));


                operationService.getServiceMeta().forEach((k, v) -> {
                    if (v.getService().equalsIgnoreCase(meta.getService())) {
                        export.add(String.format("export * from \"./%s\"", toApiFilename(formatMethodName(k))));
                    }
                });

                break;
            }
            case WS: {
                String suffix = "event";
                String fileName = "./" + toModelFilename(meta.getMethod()) + "_" + suffix;
                export.add(String.format("export * from \"%s\"", fileName));
                break;

            }
            case ENTRY: {
                List<String> serviceAliases = new LinkedList<>();
                operationService.getServiceMeta().forEach((k, v) -> {
                    if (v.getService().equalsIgnoreCase(meta.getService())) {
                        String serviceAlias = v.getSubService().toUpperCase();
                        export.add(String.format("import * as %s from \"./%s\"", serviceAlias, formatPackage(v.getSubService())));
                        serviceAliases.add(serviceAlias);
                    }
                });

                String exportService = camelize(service.toLowerCase(), CamelizeOption.UPPERCASE_FIRST_CHAR);
                export.add(String.format("const %s = \n{\n%s\n};", exportService,
                        serviceAliases.stream().map(s -> "    ..." + s).collect(Collectors.joining(",\n"))));
                export.add(String.format("export default %s;", exportService));

                break;
            }
            case WS_TEST:
            case TEST: {
                // skip
                break;
            }
            default: {
                throw new RuntimeException("unsupported mode");
            }
        }
    }


    @Override
    public OperationsMap postProcessOperationsWithModels(OperationsMap objs, List<ModelMap> allModels) {
        objs = super.postProcessOperationsWithModels(objs, allModels);

        OperationMap operationMap = objs.getOperations();

        Map<String, ImportModel> imports = new HashMap<>();

        for (CodegenOperation op : operationMap.getOperation()) {
            Meta meta = SpecificationUtil.getMeta(op.vendorExtensions);
            if (meta != null) {
                switch (modeSwitch.getMode()) {
                    case ENTRY: {
                        // api entry
                        List<Map<String, String>> entryValue = new ArrayList<>();
                        operationService.getServiceMeta().forEach((k, v) -> {
                            if (v.getService().equalsIgnoreCase(meta.getService())) {
                                Map<String, String> kv = new HashMap<>();
                                kv.put("method", formatMethodName(k));
                                kv.put("methodUppercase", camelize(formatMethodName(k)));
                                kv.put("target_service", formatService(k + "API"));
                                entryValue.add(kv);
                            }
                        });
                        Map<String, Object> apiEntryInfo = new HashMap<>();
                        apiEntryInfo.put("api_entry_name", formatService(meta.getService() + "Service"));
                        apiEntryInfo.put("api_entry_value", entryValue);
                        objs.put("api_entry", apiEntryInfo);
                        entryValue.forEach(m -> {
                            generateApiImport(meta, false, imports);
                            generateApiExport(meta, false, exports);

                        });
                        break;
                    }

                    case API:
                    case TEST: {
                        if (op.hasParams) {
                            generateApiImport(meta, true, imports);
                            generateApiExport(meta, true, exports);
                        }
                        generateApiImport(meta, false, imports);
                        generateApiExport(meta, false, exports);
                        break;
                    }
                    case WS:
                    case WS_TEST: {
                        generateApiImport(meta, false, imports);
                        generateApiExport(meta, false, exports);
                        break;
                    }
                    case TEST_TEMPLATE: {
                        String reqName = meta.getMethodServiceFmt().toLowerCase() + "Req";
                        String responseName = meta.getMethodServiceFmt().toLowerCase() + "Resp";
                        allModels.stream().filter(m -> reqName.equalsIgnoreCase((String) m.get("importPath"))).
                                forEach(m -> op.vendorExtensions.put("x-request-model", m.getModel()));
                        allModels.stream().filter(m -> responseName.equalsIgnoreCase((String) m.get("importPath"))).
                                forEach(m -> op.vendorExtensions.put("x-response-model", m.getModel()));
                        break;
                    }
                }
            }
        }

        objs.put("imports", imports.values().stream().map(ImportModel::toImport).collect(Collectors.toList()));
        return objs;
    }


    private String getInnerModelType(CodegenProperty p) {
        if (p.isArray || p.isMap) {
            return getInnerModelType(p.getItems());
        }

        if (p.isModel) {
            return p.getDataType();
        }

        return null;
    }

    @Override
    public ModelsMap postProcessModels(ModelsMap objs) {
        objs = super.postProcessModels(objs);

        Map<String, ImportModel> imports = new HashMap<>();

        List<ModelMap> models = objs.getModels();
        if (models != null) {
            for (ModelMap model : models) {
                CodegenModel codegenModel = model.getModel();
                if (codegenModel != null) {

                    imports.computeIfAbsent("class-transformer", ImportModel::new).
                            component.addAll(Arrays.asList("plainToInstance", "instanceToPlain"));

                    if (codegenModel.getVendorExtensions().containsKey("x-response-model")) {
                        imports.computeIfAbsent("class-transformer", ImportModel::new).component.add("Exclude");
                    }

                    for (CodegenProperty var : codegenModel.getVars()) {

                        String innerType = getInnerModelType(var);

                        if (innerType != null) {
                            String modelName = "./" + toModelFilename(innerType);
                            imports.computeIfAbsent(modelName, ImportModel::new).component.add(innerType);
                            imports.computeIfAbsent("class-transformer", ImportModel::new).component.add("Type");
                            var.vendorExtensions.put("x-typed", String.format("@Type(() => %s)", innerType));
                        }

                        if (var.getBaseName() != null && !var.getName().equals(var.getBaseName())) {
                            var.vendorExtensions.put("x-use-base-name", true);
                            imports.computeIfAbsent("class-transformer", ImportModel::new).component.add("Expose");
                        }

                        if (var.getVendorExtensions().containsKey("x-tag-path")) {
                            imports.computeIfAbsent("reflect-metadata", ImportModel::new);
                        }
                    }

                    codegenModel.getVendorExtensions().put("x-imports", imports.values().stream().
                            map(ImportModel::toImport).collect(Collectors.toSet()));
                }
            }
        }
        return objs;
    }

    @Override
    public Map<String, Object> postProcessSupportingFileData(Map<String, Object> objs) {
        Map<String, Object> data = super.postProcessSupportingFileData(objs);
        data.put("exports", exports);
        return data;
    }
}
