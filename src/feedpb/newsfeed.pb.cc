// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: newsfeed.proto

#define INTERNAL_SUPPRESS_PROTOBUF_FIELD_DEPRECATION
#include "newsfeed.pb.h"

#include <algorithm>

#include <google/protobuf/stubs/common.h>
#include <google/protobuf/stubs/port.h>
#include <google/protobuf/stubs/once.h>
#include <google/protobuf/io/coded_stream.h>
#include <google/protobuf/wire_format_lite_inl.h>
#include <google/protobuf/descriptor.h>
#include <google/protobuf/generated_message_reflection.h>
#include <google/protobuf/reflection_ops.h>
#include <google/protobuf/wire_format.h>
// @@protoc_insertion_point(includes)

namespace feed {
namespace v1 {

namespace {


}  // namespace


void protobuf_AssignDesc_newsfeed_2eproto() GOOGLE_ATTRIBUTE_COLD;
void protobuf_AssignDesc_newsfeed_2eproto() {
  protobuf_AddDesc_newsfeed_2eproto();
  const ::google::protobuf::FileDescriptor* file =
    ::google::protobuf::DescriptorPool::generated_pool()->FindFileByName(
      "newsfeed.proto");
  GOOGLE_CHECK(file != NULL);
}

namespace {

GOOGLE_PROTOBUF_DECLARE_ONCE(protobuf_AssignDescriptors_once_);
inline void protobuf_AssignDescriptorsOnce() {
  ::google::protobuf::GoogleOnceInit(&protobuf_AssignDescriptors_once_,
                 &protobuf_AssignDesc_newsfeed_2eproto);
}

void protobuf_RegisterTypes(const ::std::string&) GOOGLE_ATTRIBUTE_COLD;
void protobuf_RegisterTypes(const ::std::string&) {
  protobuf_AssignDescriptorsOnce();
}

}  // namespace

void protobuf_ShutdownFile_newsfeed_2eproto() {
}

void protobuf_AddDesc_newsfeed_2eproto() GOOGLE_ATTRIBUTE_COLD;
void protobuf_AddDesc_newsfeed_2eproto() {
  static bool already_here = false;
  if (already_here) return;
  already_here = true;
  GOOGLE_PROTOBUF_VERIFY_VERSION;

  ::google::api::protobuf_AddDesc_google_2fapi_2fannotations_2eproto();
  ::feed::v1::protobuf_AddDesc_common_2eproto();
  ::google::protobuf::DescriptorPool::InternalAddGeneratedFile(
    "\n\016newsfeed.proto\022\007feed.v1\032\034google/api/an"
    "notations.proto\032\014common.proto2Z\n\010NewsFee"
    "d\022N\n\003Get\022\027.feed.v1.FeedGetRequest\032\030.feed"
    ".v1.FeedGetResponse\"\024\202\323\344\223\002\016\022\014/feed/v1/ge"
    "tB\010Z\006feedpbb\006proto3", 179);
  ::google::protobuf::MessageFactory::InternalRegisterGeneratedFile(
    "newsfeed.proto", &protobuf_RegisterTypes);
  ::google::protobuf::internal::OnShutdown(&protobuf_ShutdownFile_newsfeed_2eproto);
}

// Force AddDescriptors() to be called at static initialization time.
struct StaticDescriptorInitializer_newsfeed_2eproto {
  StaticDescriptorInitializer_newsfeed_2eproto() {
    protobuf_AddDesc_newsfeed_2eproto();
  }
} static_descriptor_initializer_newsfeed_2eproto_;

// @@protoc_insertion_point(namespace_scope)

}  // namespace v1
}  // namespace feed

// @@protoc_insertion_point(global_scope)
