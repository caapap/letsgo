# ☸️ Kubernetes工程师学习路径

> 专注于容器编排和云原生技术的系统化学习指南

## 🎯 岗位技能图谱

### 核心技能（必备）
- ✅ **K8s基础**: Pod、Service、Deployment、ConfigMap
- ✅ **集群管理**: 安装部署、升级、备份恢复
- ✅ **网络**: CNI插件、Service网络、Ingress
- ✅ **存储**: PV/PVC、StorageClass、CSI
- ✅ **安全**: RBAC、NetworkPolicy、Pod安全策略

### 进阶技能（加分）
- 🌟 **自定义资源**: CRD、Operator开发
- 🌟 **监控**: Prometheus、Grafana、日志收集
- 🌟 **CI/CD**: GitOps、ArgoCD、Tekton
- 🌟 **服务网格**: Istio、Linkerd
- 🌟 **多集群**: 联邦、跨云部署

## 📚 学习路径

### 第一阶段：K8s基础（2周）
1. [容器技术基础](../foundations/container-basics.md)
2. [K8s架构原理](../specializations/k8s-architecture.md)
3. [核心资源对象](../specializations/k8s-resources.md)

### 第二阶段：集群运维（2周）
1. [集群安装部署](../specializations/k8s-installation.md)
2. [网络配置管理](../specializations/k8s-networking.md)
3. [存储解决方案](../specializations/k8s-storage.md)

### 第三阶段：高级特性（2周）
1. [自定义控制器](../specializations/k8s-controllers.md)
2. [Operator开发](../specializations/k8s-operators.md)
3. [集群安全加固](../specializations/k8s-security.md)

### 第四阶段：生产实践（1周）
1. [监控告警体系](../practices/k8s-monitoring.md)
2. [故障排查手册](../practices/k8s-troubleshooting.md)
3. [性能调优指南](../practices/k8s-performance.md)

## 🎯 面试重点

### 高频面试题
- K8s架构组件及作用
- Pod生命周期和调度
- Service网络实现原理
- 存储卷类型和使用场景
- RBAC权限控制

### 实战经验准备
- 大规模集群管理经验
- 网络故障排查案例
- 存储方案选型经验
- 监控体系建设

## 🔧 实战项目

### 入门项目
- [ ] 单节点K8s环境搭建
- [ ] 微服务应用部署
- [ ] 配置管理实践

### 进阶项目
- [ ] 多节点生产集群
- [ ] CI/CD流水线集成
- [ ] 监控告警系统

### 高级项目
- [ ] 多集群管理平台
- [ ] 自定义Operator
- [ ] 服务网格部署

## 📖 学习资源

### 官方文档
- Kubernetes官方文档
- CNCF项目文档
- 云厂商K8s服务文档

### 认证考试
- CKA (Certified Kubernetes Administrator)
- CKAD (Certified Kubernetes Application Developer)
- CKS (Certified Kubernetes Security Specialist)

### 实践环境
- Minikube本地环境
- Kind轻量级集群
- 云厂商托管K8s 