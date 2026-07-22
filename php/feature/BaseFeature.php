<?php
declare(strict_types=1);

// InfranodeOpenData SDK base feature

class InfranodeOpenDataBaseFeature
{
    public string $version;
    public string $name;
    public bool $active;

    // Positions this feature when added via the client `extend` option:
    // "__before__" / "__after__" / "__replace__" name an already-added
    // feature (mirrors the ts feature `_options`). Declared so setting it
    // on an extension instance avoids the dynamic-property deprecation.
    public ?array $_options = null;

    public function __construct()
    {
        $this->version = '0.0.1';
        $this->name = 'base';
        $this->active = true;
    }

    public function get_version(): string { return $this->version; }
    public function get_name(): string { return $this->name; }
    public function get_active(): bool { return $this->active; }

    public function init(InfranodeOpenDataContext $ctx, array $options): void {}
    public function PostConstruct(InfranodeOpenDataContext $ctx): void {}
    public function PostConstructEntity(InfranodeOpenDataContext $ctx): void {}
    public function SetData(InfranodeOpenDataContext $ctx): void {}
    public function GetData(InfranodeOpenDataContext $ctx): void {}
    public function GetMatch(InfranodeOpenDataContext $ctx): void {}
    public function SetMatch(InfranodeOpenDataContext $ctx): void {}
    public function PrePoint(InfranodeOpenDataContext $ctx): void {}
    public function PreSpec(InfranodeOpenDataContext $ctx): void {}
    public function PreRequest(InfranodeOpenDataContext $ctx): void {}
    public function PreResponse(InfranodeOpenDataContext $ctx): void {}
    public function PreResult(InfranodeOpenDataContext $ctx): void {}
    public function PreDone(InfranodeOpenDataContext $ctx): void {}
    public function PreUnexpected(InfranodeOpenDataContext $ctx): void {}
}
